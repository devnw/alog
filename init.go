package alog

import (
	"context"
	"errors"
	"io"
	"os"
	"sync"
	"time"

	"github.com/benjivesterby/validator"
)

const (
	std   = 1 << iota // use a standard logging scheme without adding delimits
	delim             // use a delimited logging scheme using a delimiter format
	json              // marshal composite structs to json for logging
)

var mutty = sync.Mutex{}

// instance is the default logger for library include, this can be replaced by
// a different default instance and is global to the library
var instance Logger

func init() {
	var err error
	if err = NewGlobal(
		context.Background(), // Default context
		std,                  // Standard logging - Strings
		"",                   // No prefix
		time.RFC3339,         // Standard time format
		time.UTC,             // UTC logging
		false,                // Debug logging disabled
		100,                  // Default buffer of 100 logs
		os.Stdout,            // Default destination: Standard Out
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
// Format: standard = 0, delim = 1, json = 2
// Prefix: the string prefix of the logs that is listed before the date
// Buffer: the buffer size of the channel for processing the logs, DEFAULT: 100
// DateFormat: the format of the date. Default: time.RFC3339
// Location: the location for logging time. Default: time.UTC
// Debug: enable debug logging: Default: false
func NewGlobal(ctx context.Context, format int, prefix string, dateformat string, location *time.Location, debug bool, buffer int, out ...io.Writer) (err error) {
	var newLogger Logger
	if newLogger, err = New(ctx, format, prefix, dateformat, location, debug, buffer, out...); err == nil {
		err = setGlobal(newLogger)
	}

	return err
}

// New creates a new logger using the information passed in to setup
// the logging configuration rather than using the standard Stdout logger
// that is initialized automatically
// Format: standard = 0, delim = 1, json = 2
// Prefix: the string prefix of the logs that is listed before the date
// Buffer: the buffer size of the channel for processing the logs, DEFAULT: 100
// DateFormat: the format of the date. Default: time.RFC3339
// Location: the location for logging time. Default: time.UTC
// Debug: enable debug logging: Default: false
func New(ctx context.Context, format int, prefix string, dateformat string, location *time.Location, debug bool, buffer int, out ...io.Writer) (logger Logger, err error) {

	if ctx == nil {
		ctx = context.Background()
	}

	if format >= 0 {
		if buffer >= 0 {
			if len(out) >= 1 {

				// Setup the logger context
				var cancel context.CancelFunc
				ctx, cancel = context.WithCancel(ctx)

				// Initialize the logger struct
				a := &alog{
					ctx:        ctx,
					cancel:     cancel,
					outputs:    out,
					format:     format,
					prefix:     prefix,
					location:   time.UTC,
					dateformat: time.RFC3339,
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
	} else {
		// TODO:
	}

	return logger, err
}
