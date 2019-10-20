package alog

import (
	"context"
	"io"
	"os"
	"time"
)

const (
	std   = 1 << iota // use a standard logging scheme without adding delimits
	delim             // use a delimited logging scheme using a delimiter format
	json              // marshal composite structs to json for logging
)

// DateFormat is the date format that is used in the logging. Default: RFC3339
var DateFormat = time.RFC3339

// Location is the timezone that is used for logging. Default: UTC
var Location = time.UTC

// Instance is the default logger for library include, this can be replaced by
// a different default instance and is global to the library
var Instance, _ = New(context.Background(), std, "", os.Stdout)

// New creates a new logger using the information passed in to setup
// the logging configuration rather than using the standard Stdout logger
// that is initialized automatically
func New(ctx context.Context, format int, prefix string, out ...io.Writer) (Logger, error) {
	var err error

	// Setup the logger context
	var cancel context.CancelFunc
	ctx, cancel = context.WithCancel(ctx)

	// Initialize the logger struct
	logger := &alog{
		ctx:     ctx,
		cancel:  cancel,
		outputs: out,
		format:  format,
		prefix:  prefix,
	}

	// TODO: initialize the go routines for reading the logs

	return logger, err
}
