package alog

import (
	"context"
	"sync"
	"time"

	"github.com/benjivesterby/validator"
	"github.com/pkg/errors"
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

const (
	// DEFAULTBUFFER is the global default for buffer size on the async logger
	DEFAULTBUFFER = 0

	// DEFAULTTIMEFORMAT is the global default time format used by the async logger
	DEFAULTTIMEFORMAT = time.RFC3339
)

var mutty = sync.Mutex{}

// instance is the default logger for library include, this can be replaced by
// a different default instance and is global to the library
var instance Logger

func init() {
	var err error

	if err = Global(
		context.Background(), // Default context
		"",                   // No prefix
		DEFAULTTIMEFORMAT,    // Standard time format
		time.UTC,             // UTC logging
		false,                // Debug logging disabled
		DEFAULTBUFFER,        // Default buffer of 100 logs
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

// Global instantiates a new logger using the passed in parameters and overwrites
// the package global instance of the logger
func Global(ctx context.Context, prefix string, dateformat string, location *time.Location, debug bool, buffer int, destinations ...Destination) (err error) {
	var newLogger Logger
	if newLogger, err = New(ctx, prefix, dateformat, location, debug, buffer, destinations...); err == nil {
		err = setGlobal(newLogger)
	}

	return err
}

// New creates a new logger using the information passed in to setup
// the logging configuration rather than using the standard Stdout logger
// that is initialized automatically
func New(ctx context.Context, prefix string, dateformat string, location *time.Location, debug bool, buffer int, destinations ...Destination) (logger Logger, err error) {

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
			location:     time.UTC,
			dateformat:   time.RFC3339,
			logdebug:     debug,
			buffer:       buffer,
		}

		// initialize the go routines for reading the logs
		if err = a.init(); err == nil {
			logger = a
		} else {
			// cancel the context and error out the logger creation
			defer a.cancel()
			err = errors.Errorf("error occurred while initializing logger | %s", err)
		}
	} else {
		err = errors.New("logger requires at least one destination")
	}

	return logger, err
}
