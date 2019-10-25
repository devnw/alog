package alog

import (
	"context"
	"io"
	"time"

	"github.com/benjivesterby/validator"
)

// TODO: Setup so that new destinations can be added at runtime (chan Dest)
// TODO: flag for stack traces on logs with errors?

type alog struct {
	ctx          context.Context
	cancel       context.CancelFunc
	destinations []Dest

	// location is the timezone that is used for logging. Default: UTC
	location *time.Location

	// dateformat is the date format that is used in the logging. Default: RFC3339
	dateformat string
	prefix     string
	logdebug   bool

	// The channels which will have logs sent and received on
	info  chan log
	debug chan log
	warn  chan log
	err   chan log
	crit  chan log
	fatal chan log
}

// init starts up the go routines for receiving and publishing logs
// to the available io.Writers
func (l *alog) init() (err error) {
	return err
}

// Print creates informational logs based on the inputs
func (l *alog) Print(v ...interface{}) {

}

// Println prints the data coming in as an informational log on individual lines
func (l *alog) Println(v ...interface{}) {

}

// Printf creates an informational log using the format and values
func (l *alog) Printf(format string, v ...interface{}) {

}

// Debug creates debugging logs based on the inputs
func (l *alog) Debug(v ...interface{}) {
	if l.logdebug {

	}
}

// Debugln prints the data coming in as a debug log on individual lines
func (l *alog) Debugln(v ...interface{}) {
	if l.logdebug {

	}
}

// Debugf creates an debugging log using the format and values
func (l *alog) Debugf(format string, v ...interface{}) {
	if l.logdebug {

	}
}

// Warn creates a warning log using the error passed in along with the
// values passed in
func (l *alog) Warn(err error, v ...interface{}) {

}

// Warnln creates a warning log using the error and values passed in.
// Each error and value is printed on a different line
func (l *alog) Warnln(err error, v ...interface{}) {

}

// Warnf creates a warning log using the error passed in, along with the string
// formatting and values
func (l *alog) Warnf(err error, format string, v ...interface{}) {

}

// Error creates an error log using the error and other values passed in
func (l *alog) Error(err error, v ...interface{}) {

}

// Errorln creates error logs using the error and other values passed in.
// Each error and value is printed on a different line
func (l *alog) Errorln(err error, v ...interface{}) {

}

// Errorf creates an error log using the error passed in, along with the string
// formatting and values
func (l *alog) Errorf(err error, format string, v ...interface{}) {

}

// Crit creates critical logs using the error and other values passed in
func (l *alog) Crit(err error, v ...interface{}) {

}

// Critln creates critical logs using the error and other values passed in.
// Each error and value is printed on a different line
func (l *alog) Critln(err error, v ...interface{}) {

}

// Critf creates a critical log using the error passed in, along with the string
// formatting and values
func (l *alog) Critf(err error, format string, v ...interface{}) {

}

// Fatal creates a fatal log using the error and values passed into the method
// After logging the fatal log the Fatal method throws a panic to crash the application
func (l *alog) Fatal(err error, v ...interface{}) {

	// TODO: Update panic to include information about the fatal, as well as stack trace information
	panic(err)
}

// Fatalln creates fatal logs using the error and other values passed in.
// Each error and value is printed on a different line
// After logging the fatal log the Fatalln method throws a panic to crash the application
func (l *alog) Fatalln(err error, v ...interface{}) {

	// TODO: Update panic to include information about the fatal, as well as stack trace information
	panic(err)
}

// Fatalf creates an error log using the error passed in, along with the string
// formatting and values
// After logging the fatal log the Fatalf method throws a panic to crash the application
func (l *alog) Fatalf(err error, format string, v ...interface{}) {

	// TODO: Update panic to include information about the fatal, as well as stack trace information
	panic(err)
}

// Custom creates a custom log using the error and values passed into the method
func (l *alog) Custom(ltype string, err error, v ...interface{}) {
}

// Customln creates custom logs using the error and other values passed in.
// Each error and value is printed on a different line
func (l *alog) Customln(ltype string, err error, v ...interface{}) {
}

// Customf creates a custom log using the error passed in, along with the string
// formatting and values
func (l *alog) Customf(ltype string, err error, format string, v ...interface{}) {
}

// AddOutput adds an additional logging source to the logger which
// will be added to the different logging outputs for this logger
func (l *alog) AddOutput(out io.Writer) {

}

// Close cancels the context of the logger internally and breaks out of
// any logging activity. This should always be called in a defer at the top
// level where the logger is initialized to ensure proper closure
func (l *alog) Close() {
	if validator.IsValid(l) {

		// cancel the context of the logger
		l.cancel()
	}
}

// Validate checks the validity and health of the logger
// to ensure that it can properly log
func (l *alog) Validate() (valid bool) {
	if l != nil && l.ctx != nil && l.cancel != nil {
		// TODO: ensure there is at least one io.Writer registered and
		// that the health checks are passing
		valid = true
	}

	return valid
}
