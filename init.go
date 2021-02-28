// Copyright © 2019 Developer Network, LLC
//
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of this source code package.

package alog

import (
	"context"
	"errors"
	"sync"
	"time"

	"devnw.com/validator"
)

// LogFmt is the logging format
type LogFmt int8

const (
	// STD uses a standard logging scheme without adding delimits
	STD LogFmt = 1 << iota

	// JSON marshals composite structs to json for logging
	JSON
)

// LogLevel is the logging level for the logger
type LogLevel int16

func (l LogLevel) String() string {
	t, ok := types[l]
	if !ok {
		// Default to INFO type for unknown log types
		return types[INFO]
	}

	return t
}

// Log switches for sources setup using bitwise comparison
const (
	// INFO is the flag for logging informational logs on a destination
	INFO LogLevel = 1 << iota

	// DEBUG is the flag for logging debugging logs on a destination
	DEBUG

	// TRACE is the flag for logging trace logs on a destination
	TRACE

	// WARN is the flag for logging warning logs on a destination
	WARN

	// ERROR is the flag for logging error logs on a destination
	ERROR

	// CRIT is the flag for logging critical logs on a destination
	CRIT

	// FATAL is the flag for logging fatal logs on a destination
	FATAL

	// CUSTOM is the flag to indicate that a custom log type is being passed
	// instead of one of the built in types
	CUSTOM
)

var types = map[LogLevel]string{
	INFO:   "INFO",
	DEBUG:  "DEBUG",
	TRACE:  "TRACE",
	WARN:   "WARN",
	ERROR:  "ERROR",
	CRIT:   "CRITICAL",
	FATAL:  "FATAL",
	CUSTOM: "CUSTOM",
}

const (
	// DEFAULTBUFFER is the global default for buffer size on the async logger
	DEFAULTBUFFER = 0

	// DEFAULTTIMEFORMAT is the global default time format used by the async logger
	DEFAULTTIMEFORMAT = time.RFC3339
)

// global is the default logger for library include, this can be replaced by
// a different default global and is global to the library
var global Logger
var globalMu = sync.Mutex{}

func init() {
	_ = Global(
		context.Background(), // Default context
		"",                   // No prefix
		DEFAULTTIMEFORMAT,    // Standard time format
		time.UTC,             // UTC logging
		DEFAULTBUFFER,        // Default buffer of 100 logs
		Standards()...,       // Default destinations
	)
}

// setGlobal overrides the default logger using the passed in logger for the library
// to facilitate simplified global logging
func setGlobal(logger Logger) (err error) {
	if validator.Valid(logger) {
		globalMu.Lock()
		defer globalMu.Unlock()

		// Close the logger instance and replace it
		if global != nil {
			global.Close()
		}

		global = logger
	} else {
		err = errors.New("invalid logger")
	}

	return err
}

// Global instantiates a new logger using the passed in parameters and
// overwrites the package global instance of the logger
func Global(
	ctx context.Context,
	prefix string,
	dateformat string,
	location *time.Location,
	buffer int,
	destinations ...Destination,
) error {
	newLogger, err := New(
		ctx,
		prefix,
		dateformat,
		location,
		buffer,
		destinations...,
	)

	if err != nil {
		return err
	}

	return setGlobal(newLogger)
}

// New creates a new logger using the information passed in to setup
// the logging configuration rather than using the standard Stdout logger
// that is initialized automatically
func New(
	ctx context.Context,
	prefix string,
	dateformat string,
	location *time.Location,
	buffer int,
	destinations ...Destination,
) (logger Logger, err error) {

	if ctx == nil {
		ctx = context.Background()
	}

	if location == nil {
		location = time.UTC
	}

	if len(dateformat) == 0 {
		dateformat = DEFAULTTIMEFORMAT
	}

	// Default an invalid buffer to the default
	if buffer < 0 {
		buffer = DEFAULTBUFFER
	}

	if len(destinations) >= 1 {

		// Setup the logger context
		var cancel context.CancelFunc
		ctx, cancel = context.WithCancel(ctx)

		// Initialize the logger struct
		a := &alog{
			ctx:          ctx,
			cancel:       cancel,
			destinations: destinations,
			prefix:       prefix,
			location:     location,
			dateformat:   time.RFC3339,
			buffer:       buffer,
			mutty:        sync.RWMutex{},
			cleaned:      make(chan bool),
			out: map[LogLevel][]chan<- log{
				INFO:   make([]chan<- log, 0),
				DEBUG:  make([]chan<- log, 0),
				TRACE:  make([]chan<- log, 0),
				WARN:   make([]chan<- log, 0),
				ERROR:  make([]chan<- log, 0),
				CRIT:   make([]chan<- log, 0),
				FATAL:  make([]chan<- log, 0),
				CUSTOM: make([]chan<- log, 0),
			},
		}

		// initialize the go routines for reading the logs
		a.init()
		logger = a
	} else {
		err = errors.New("logger requires at least one destination")
	}

	return logger, err
}
