package alog

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/benjivesterby/validator"
)

const (
	// STD uses a standard logging scheme without adding delimits
	STD = 0

	// DELIM uses a delimited logging scheme using a delimiter format. DEFAULT: csv
	DELIM = 1

	// JSON marshals composite structs to json for logging
	JSON = 2
)

// Log switches for sources setup using bitwise comparision
const (
	// INFO is the flag for logging informational logs on a destination
	INFO = 1 << iota

	// DEBUG is the flag for logging debugging logs on a destination
	DEBUG

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

var mutty = sync.Mutex{}

// instance is the default logger for library include, this can be replaced by
// a different default instance and is global to the library
var instance Logger

func init() {
	var err error

	if err = NewGlobal(
		context.Background(), // Default context
		"",                   // No prefix
		time.RFC3339,         // Standard time format
		time.UTC,             // UTC logging
		false,                // Debug logging disabled
		100,                  // Default buffer of 100 logs
		Standards()...,       // Default destinations
	); err != nil {

		// Panic if the initialization fails
		panic(err)
	}
}

// setGlobal overrides the default logger using the passed in logger for the library
// to facilitate simplified global logging
func setGlobal(logger Logger) (err error) {
	if validator.IsValid(logger) {
		mutty.Lock()
		defer mutty.Unlock()

		// Close the logger instance and replace it
		instance.Close()
		instance = logger
	} else {
		err = errors.New("invalid logger")
	}

	return err
}

// NewGlobal instantiates a new logger using the passed in parameters and overwrites
// the package global instance of the logger
// Prefix: the string prefix of the logs that is listed before the date
// Buffer: the buffer size of the channel for processing the logs, DEFAULT: 100
// DateFormat: the format of the date. Default: time.RFC3339
// Location: the location for logging time. Default: time.UTC
// Debug: enable debug logging: Default: false
func NewGlobal(ctx context.Context, prefix string, dateformat string, location *time.Location, debug bool, buffer int, destinations ...Dest) (err error) {
	var newLogger Logger
	if newLogger, err = New(ctx, prefix, dateformat, location, debug, buffer, destinations...); err == nil {
		err = setGlobal(newLogger)
	}

	return err
}

// New creates a new logger using the information passed in to setup
// the logging configuration rather than using the standard Stdout logger
// that is initialized automatically
// Prefix: the string prefix of the logs that is listed before the date
// Buffer: the buffer size of the channel for processing the logs, DEFAULT: 100
// DateFormat: the format of the date. Default: time.RFC3339
// Location: the location for logging time. Default: time.UTC
// Debug: enable debug logging: Default: false
func New(ctx context.Context, prefix string, dateformat string, location *time.Location, debug bool, buffer int, destinations ...Dest) (logger Logger, err error) {

	if ctx == nil {
		ctx = context.Background()
	}

	if location == nil {
		location = time.UTC
	}

	if len(dateformat) == 0 {
		dateformat = time.RFC3339
	}

	if buffer >= 0 {
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
				location:     time.UTC,
				dateformat:   time.RFC3339,
				logdebug:     debug,
			}

			// initialize the go routines for reading the logs
			if err = a.init(); err == nil {
				logger = a
			}
		} else {
			// TODO:
		}
	} else {
		// TODO:
	}

	return logger, err
}
