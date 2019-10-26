package alog

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/benjivesterby/validator"
)

// TODO: Setup so that new destinations can be added at runtime (chan Dest)
// TODO: flag for stack traces on logs with errors?

type alog struct {
	ctx          context.Context
	cancel       context.CancelFunc
	destinations []Destination

	// location is the timezone that is used for logging. Default: UTC
	location *time.Location

	// dateformat is the date format that is used in the logging. Default: RFC3339
	dateformat string
	prefix     string
	logdebug   bool
	buffer     int

	// The channels which will have logs sent and received on
	info      chan log
	infodests []chan<- log

	debug      chan log
	debugdests []chan<- log

	warn      chan log
	warndests []chan<- log

	err      chan log
	errdests []chan<- log

	crit      chan log
	critdests []chan<- log

	fatal      chan log
	fataldests []chan<- log

	custom      chan log
	customdests []chan<- log
}

// init starts up the go routines for receiving and publishing logs
// to the available io.Writers
func (l *alog) init() (err error) {

	for _, dest := range l.destinations {
		if dest.Types&INFO > 0 {
			l.infodests = append(l.infodests, l.listen(l.ctx, dest))
		}

		if dest.Types&DEBUG > 0 {
			l.debugdests = append(l.debugdests, l.listen(l.ctx, dest))
		}

		if dest.Types&WARN > 0 {
			l.warndests = append(l.warndests, l.listen(l.ctx, dest))
		}

		if dest.Types&ERROR > 0 {
			l.errdests = append(l.errdests, l.listen(l.ctx, dest))
		}

		if dest.Types&CRIT > 0 {
			l.critdests = append(l.critdests, l.listen(l.ctx, dest))
		}

		if dest.Types&FATAL > 0 {
			l.fataldests = append(l.fataldests, l.listen(l.ctx, dest))
		}

		if dest.Types&CUSTOM > 0 {
			l.customdests = append(l.customdests, l.listen(l.ctx, dest))
		}
	}

	return err
}

func (l *alog) listen(ctx context.Context, destination Destination) chan<- log {
	logs := make(chan log)

	go func(ctx context.Context, logs chan log, destination Destination) {
		// TODO: handle panic
		defer close(logs)

		for {
			select {
			case <-ctx.Done():
				// TODO: setup to close the destination if it has a close method
				return
			case log, ok := <-logs:
				if ok {
					if validator.IsValid(log) {
						// TODO: correct this
						fmt.Fprint(destination.Writer, log)
					} else {
						// TODO:
					}
				} else {
					return
				}
			}
		}
	}(ctx, logs, destination)

	return logs
}

// send is used to create a go routine thread for fanning out specific
// log types to each of the destinations
func (l *alog) send(value log) {
	// TODO: Handle panic here

	// Loop over the destinations for this logtype and push onto the
	// log channels for each destination
	for _, destination := range l.getd(value.logtype) {

		// Push the log onto the destination channel
		select {
		case <-l.ctx.Done():
			return
		case destination <- value:
		}
	}
}

func (l *alog) getd(logtype int8) []chan<- log {
	destinations := l.infodests

	if logtype&DEBUG > 0 {
		destinations = l.debugdests
	} else if logtype&WARN > 0 {
		destinations = l.warndests
	} else if logtype&ERROR > 0 {
		destinations = l.errdests
	} else if logtype&CRIT > 0 {
		destinations = l.critdests
	} else if logtype&FATAL > 0 {
		destinations = l.fataldests
	} else if logtype&CUSTOM > 0 {
		destinations = l.customdests
	}

	return destinations
}

func (l *alog) buildlog(logtype int8, custom string, err error, format *string, v ...interface{}) (newlog log) {

	values := v
	if format != nil {
		values = []interface{}{fmt.Sprintf(*format, v...)}
	}

	newlog = log{
		logger:     l,
		logtype:    logtype,
		customtype: custom,
		timestamp:  time.Now(),
		err:        err,
		values:     values,
	}

	return newlog
}

// Print creates informational logs based on the inputs
func (l *alog) Print(v ...interface{}) {
	l.send(l.buildlog(INFO, "", nil, nil, v...))
}

// Println prints the data coming in as an informational log on individual lines
func (l *alog) Println(v ...interface{}) {
	for _, value := range v {
		l.send(l.buildlog(INFO, "", nil, nil, value))
	}
}

// Printf creates an informational log using the format and values
func (l *alog) Printf(format string, v ...interface{}) {
	l.send(l.buildlog(INFO, "", nil, &format, v...))
}

// Debug creates debugging logs based on the inputs
func (l *alog) Debug(v ...interface{}) {
	if l.logdebug {
		l.send(l.buildlog(DEBUG, "", nil, nil, v...))
	}
}

// Debugln prints the data coming in as a debug log on individual lines
func (l *alog) Debugln(v ...interface{}) {
	if l.logdebug {
		for _, value := range v {
			l.send(l.buildlog(DEBUG, "", nil, nil, value))
		}
	}
}

// Debugf creates an debugging log using the format and values
func (l *alog) Debugf(format string, v ...interface{}) {
	if l.logdebug {
		l.send(l.buildlog(DEBUG, "", nil, &format, v...))
	}
}

// Warn creates a warning log using the error passed in along with the
// values passed in
func (l *alog) Warn(err error, v ...interface{}) {
	l.send(l.buildlog(WARN, "", nil, nil, v...))
}

// Warnln creates a warning log using the error and values passed in.
// Each error and value is printed on a different line
func (l *alog) Warnln(err error, v ...interface{}) {
	for _, value := range v {
		l.send(l.buildlog(WARN, "", nil, nil, value))
	}
}

// Warnf creates a warning log using the error passed in, along with the string
// formatting and values
func (l *alog) Warnf(err error, format string, v ...interface{}) {
	l.send(l.buildlog(WARN, "", nil, &format, v...))
}

// Error creates an error log using the error and other values passed in
func (l *alog) Error(err error, v ...interface{}) {
	l.send(l.buildlog(ERROR, "", nil, nil, v...))
}

// Errorln creates error logs using the error and other values passed in.
// Each error and value is printed on a different line
func (l *alog) Errorln(err error, v ...interface{}) {
	for _, value := range v {
		l.send(l.buildlog(ERROR, "", nil, nil, value))
	}
}

// Errorf creates an error log using the error passed in, along with the string
// formatting and values
func (l *alog) Errorf(err error, format string, v ...interface{}) {
	l.send(l.buildlog(ERROR, "", nil, &format, v...))
}

// Crit creates critical logs using the error and other values passed in
func (l *alog) Crit(err error, v ...interface{}) {
	l.send(l.buildlog(CRIT, "", nil, nil, v...))
}

// Critln creates critical logs using the error and other values passed in.
// Each error and value is printed on a different line
func (l *alog) Critln(err error, v ...interface{}) {
	for _, value := range v {
		l.send(l.buildlog(CRIT, "", nil, nil, value))
	}
}

// Critf creates a critical log using the error passed in, along with the string
// formatting and values
func (l *alog) Critf(err error, format string, v ...interface{}) {
	l.send(l.buildlog(CRIT, "", nil, &format, v...))
}

// Fatal creates a fatal log using the error and values passed into the method
// After logging the fatal log the Fatal method throws a panic to crash the application
func (l *alog) Fatal(err error, v ...interface{}) {
	l.send(l.buildlog(FATAL, "", nil, nil, v...))

	// TODO: Update panic to include information about the fatal, as well as stack trace information
	panic(err)
}

// Fatalln creates fatal logs using the error and other values passed in.
// Each error and value is printed on a different line
// After logging the fatal log the Fatalln method throws a panic to crash the application
func (l *alog) Fatalln(err error, v ...interface{}) {
	for _, value := range v {
		l.send(l.buildlog(FATAL, "", nil, nil, value))
	}

	// TODO: Update panic to include information about the fatal, as well as stack trace information
	panic(err)
}

// Fatalf creates an error log using the error passed in, along with the string
// formatting and values
// After logging the fatal log the Fatalf method throws a panic to crash the application
func (l *alog) Fatalf(err error, format string, v ...interface{}) {
	l.send(l.buildlog(FATAL, "", nil, &format, v...))

	// TODO: Update panic to include information about the fatal, as well as stack trace information
	panic(err)
}

// Custom creates a custom log using the error and values passed into the method
func (l *alog) Custom(ltype string, err error, v ...interface{}) {
	l.send(l.buildlog(CUSTOM, ltype, nil, nil, v...))
}

// Customln creates custom logs using the error and other values passed in.
// Each error and value is printed on a different line
func (l *alog) Customln(ltype string, err error, v ...interface{}) {
	for _, value := range v {
		l.send(l.buildlog(CUSTOM, ltype, nil, nil, value))
	}
}

// Customf creates a custom log using the error passed in, along with the string
// formatting and values
func (l *alog) Customf(ltype string, err error, format string, v ...interface{}) {
	l.send(l.buildlog(CUSTOM, ltype, nil, &format, v...))
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
