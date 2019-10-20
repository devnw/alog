package alog

import (
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
var Instance = New(std, "", os.Stdout)

// New creates a new logger using the information passed in to setup
// the logging configuration rather than using the standard Stdout logger
// that is initialized automatically
func New(format int, prefix string, out ...io.Writer) Logger {
	return &alog{}
}
