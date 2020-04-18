// Package alog is an implementation of a fully concurrent non-IO blocking
// log library. Internally messages are routed through a number of subscribed
// destinations consisting of destination log levels, timezone formatting
// strings and an io.Writer for the destination to be written to.
//
// The globaly available functions utilize a singleton implementation of a
// global logger that is initialized when the package is loaded. This global
// logger utilizes os.Stdout for INFO, DEBUG, TRACE, WARN, CUSTOM logs, and
// os.Stderr for ERROR, CRITICAL, FATAL as defined in the method Standard()
//
// The global logger singleton can be overridden using the "Global" method
// passing in the proper parameters as required by your use case.
//
// If you prefer to use a non-global singleton logger you can use the "New"
// method to create a new logger which implements the alog packages "Logger"
// interface which can be passed throughout an application as a logger.
//
// Each log level has an associated function:
//
// INFO - Print, Printf, Println, Printc
//
// DEBUG - Debug, Debugf, Debugln, Debugc
//
// TRACE - Trace, Tracef, Traceln, Tracec
//
// WARN - Warn, Warnf, Warnln, Warnc
//
// ERROR - Error, Errorf, Errorln, Errorc
//
// CRITICAL - Crit, Critf, Critln, Critc
//
// FATAL - Fatal, Fatalf, Fatalln, Fatalc
//
// CUSTOM Level - Custom, Customf, Customln, Customc
//
// c(channel) functions are special methods in the logger that rather than being called
// directly are passed a channel which receives logs into the logger in a
// concurrent manner without direct calls to the `c(channel) function` again afterwards.
// These methods are particularly useful for applications where the overhead of
// a logger in concurrent actions is a particular nuisance. For these cases the
// c(channel) functions are a perfect use case. Each log level supports it's own
// c(channel) function.
//
// alog currently support two log formatting types
// * STD => Formats your logs using a "PREFIX DATE [LEVEL] Log | err: Error"
// * JSON => Formats your logs using JSON in the following form
// `{"prefix":"PREFIX","type":"ERROR","timestamp":"2020-03-01T16:21:28-06:00","error":"ERROR","messages":["ERROR"]}`
package alog

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/devnw/validator"
)

// TODO: Setup so that new destinations can be added at runtime (chan Dest)
// TODO: flag for stack traces on logs with errors?

// streamlog is a constant for the log value of a streaming log when an error type is sent
const streamlog = "stream log"

type alog struct {
	ctx          context.Context
	cancel       context.CancelFunc
	destinations []Destination

	// location is the timezone that is used for logging. Default: UTC
	location *time.Location

	// dateformat is the date format that is used in the logging. Default: RFC3339
	dateformat string
	prefix     string
	buffer     int

	mutty sync.RWMutex

	// The channels which will have logs sent and received on
	infodests   []chan<- log
	debugdests  []chan<- log
	tracedests  []chan<- log
	warndests   []chan<- log
	errdests    []chan<- log
	critdests   []chan<- log
	fataldests  []chan<- log
	customdests []chan<- log

	// Indicates that all logs have been cleared to the respective
	// destinations during a close
	cleaned chan bool
}

func (l *alog) cleanup() {
	<-l.ctx.Done()

	// Lock the destinations
	l.mutty.Lock()
	defer l.mutty.Unlock()
	defer close(l.cleaned)

	// Cleanup the destinations
	l.clean(INFO)
	l.clean(DEBUG)
	l.clean(TRACE)
	l.clean(WARN)
	l.clean(ERROR)
	l.clean(CRIT)
	l.clean(FATAL)
	l.clean(CUSTOM)
}

func (l *alog) clean(logtype LogLevel) {
	// Loop over the destinations for this logtype and close the channels
	for _, destination := range l.getd(logtype) {
		close(destination)
	}
}

// init starts up the go routines for receiving and publishing logs
// to the available io.Writers
func (l *alog) init() {

	// Startup the cleanup go routine to monitor for the closed context switch
	go l.cleanup()

	for _, dest := range l.destinations {
		if dest.Types&INFO > 0 {

			l.infodests = append(l.infodests, l.listen(l.ctx, dest))
		}

		if dest.Types&DEBUG > 0 {

			l.debugdests = append(l.debugdests, l.listen(l.ctx, dest))
		}

		if dest.Types&TRACE > 0 {

			l.tracedests = append(l.tracedests, l.listen(l.ctx, dest))
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
}

func (l *alog) listen(ctx context.Context, destination Destination) chan<- log {
	logs := make(chan log)

	go func(ctx context.Context, logs <-chan log, destination Destination) {
		// TODO: handle panic

		for {
			select {
			case <-ctx.Done():
				// TODO: setup to close the destination if it has a close method
				return
			case l, ok := <-logs:
				if ok {
					if validator.Valid(l) {
						var message string

						switch destination.Format {
						case JSON:
							if msg, err := json.Marshal(l); err == nil {

								// Add a newline to each json log for readability
								message = string(msg) + "\n"
							} else {
								// TODO: panic?
								panic("error marshalling JSON")
							}
						default:
							message = l.String()
						}

						if _, err := destination.Writer.Write([]byte(message)); err != nil {
							panic("error writing to destination")
						}
					} else {
						panic("invalid log at destination")
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
func (l *alog) send(ctx context.Context, value log) {
	// TODO: Handle panic here

	// Break out in the event that the context has been cancelled
	select {
	case <-ctx.Done():
	default:

		// Lock reads here while pulling channels
		l.mutty.RLock()
		defer l.mutty.RUnlock()

		dests := l.getd(value.logtype)

		// Loop over the destinations for this logtype and push onto the
		// log channels for each destination
		for _, destination := range dests {

			// Push the log onto the destination channel
			select {
			case <-ctx.Done():
				return
			case destination <- value:
			}
		}
	}
}

func (l *alog) getd(level LogLevel) []chan<- log {
	destinations := l.infodests

	if level&DEBUG > 0 {
		destinations = l.debugdests
	} else if level&TRACE > 0 {
		destinations = l.tracedests
	} else if level&WARN > 0 {
		destinations = l.warndests
	} else if level&ERROR > 0 {
		destinations = l.errdests
	} else if level&CRIT > 0 {
		destinations = l.critdests
	} else if level&FATAL > 0 {
		destinations = l.fataldests
	} else if level&CUSTOM > 0 {
		destinations = l.customdests
	}

	return destinations
}

func (l *alog) buildlog(logtype LogLevel, custom string, err error, format *string, time time.Time, v ...interface{}) (newlog log) {

	values := v
	if format != nil {
		values = []interface{}{fmt.Sprintf(*format, v...)}
	}

	newlog = log{
		logger:     l,
		logtype:    logtype,
		customtype: custom,
		timestamp:  time,
		err:        err,
		values:     values,
	}

	return newlog
}

func (l *alog) clog(ctx context.Context, v <-chan interface{}, level LogLevel, custom string) {

	go func(ctx context.Context, v <-chan interface{}, level LogLevel, custom string) {
		// TODO: handle panic
		for {
			select {
			case <-ctx.Done():
				return
			case value, ok := <-v:
				if ok {
					switch t := value.(type) {
					case error:
						l.send(ctx, l.buildlog(level, custom, t, nil, time.Now(), streamlog))
					default:
						l.send(ctx, l.buildlog(level, custom, nil, nil, time.Now(), t))
					}
				} else {
					return
				}
			}
		}
	}(ctx, v, level, custom)
}

// Printc creates informational logs based on the data coming from the
// concurrency channel that is passed in for processing
func (l *alog) Printc(ctx context.Context, v <-chan interface{}) {
	l.clog(ctx, v, INFO, "")
}

// Print creates informational logs based on the inputs
func (l *alog) Print(v ...interface{}) {
	t := time.Now()
	go l.send(l.ctx, l.buildlog(INFO, "", nil, nil, t, v...))
}

// Println prints the data coming in as an informational log on individual lines
func (l *alog) Println(v ...interface{}) {
	t := time.Now()
	go func(v ...interface{}) {
		if len(v) > 0 {

			for _, value := range v {
				l.send(l.ctx, l.buildlog(INFO, "", nil, nil, t, value))
			}
		} else {
			l.send(l.ctx, l.buildlog(WARN, "", nil, nil, t, "empty log value passed"))
		}
	}(v...)
}

// Printf creates an informational log using the format and values
func (l *alog) Printf(format string, v ...interface{}) {
	t := time.Now()
	go l.send(l.ctx, l.buildlog(INFO, "", nil, &format, t, v...))
}

// Debugc creates debug logs based on the data coming from the
// concurrency channel that is passed in for processing
func (l *alog) Debugc(ctx context.Context, v <-chan interface{}) {
	l.clog(ctx, v, DEBUG, "")
}

// Debug creates debugging logs based on the inputs
func (l *alog) Debug(err error, v ...interface{}) {
	t := time.Now()
	go l.send(l.ctx, l.buildlog(DEBUG, "", err, nil, t, v...))
}

// Debugln prints the data coming in as a debug log on individual lines
func (l *alog) Debugln(err error, v ...interface{}) {
	t := time.Now()
	go func(v ...interface{}) {
		if len(v) > 0 {

			for _, value := range v {
				l.send(l.ctx, l.buildlog(DEBUG, "", err, nil, t, value))
			}
		} else {
			l.send(l.ctx, l.buildlog(WARN, "", nil, nil, t, "empty log value passed"))
		}
	}(v...)
}

// Debugf creates an debugging log using the format and values
func (l *alog) Debugf(err error, format string, v ...interface{}) {
	t := time.Now()
	go l.send(l.ctx, l.buildlog(DEBUG, "", err, &format, t, v...))
}

// Tracec creates trace logs based on the data coming from the
// concurrency channel that is passed in for processing
func (l *alog) Tracec(ctx context.Context, v <-chan interface{}) {
	l.clog(ctx, v, TRACE, "")
}

// Trace creates trace logs based on the inputs
func (l *alog) Trace(err error, v ...interface{}) {
	t := time.Now()
	go l.send(l.ctx, l.buildlog(TRACE, "", err, nil, t, v...))
}

// Traceln prints the data coming in as a trace log on individual lines
func (l *alog) Traceln(err error, v ...interface{}) {
	t := time.Now()
	go func(v ...interface{}) {
		if len(v) > 0 {

			for _, value := range v {
				l.send(l.ctx, l.buildlog(TRACE, "", err, nil, t, value))
			}
		} else {
			l.send(l.ctx, l.buildlog(WARN, "", nil, nil, t, "empty log value passed"))
		}
	}(v...)
}

// Tracef creates an trace log using the format and values
func (l *alog) Tracef(err error, format string, v ...interface{}) {
	t := time.Now()
	go l.send(l.ctx, l.buildlog(TRACE, "", err, &format, t, v...))
}

// Warnc creates warning logs based on the data coming from the
// concurrency channel that is passed in for processing
func (l *alog) Warnc(ctx context.Context, v <-chan interface{}) {
	l.clog(ctx, v, WARN, "")
}

// Warn creates a warning log using the error passed in along with the
// values passed in
func (l *alog) Warn(err error, v ...interface{}) {
	t := time.Now()
	go l.send(l.ctx, l.buildlog(WARN, "", err, nil, t, v...))
}

// Warnln creates a warning log using the error and values passed in.
// Each error and value is printed on a different line
func (l *alog) Warnln(err error, v ...interface{}) {
	t := time.Now()
	go func(v ...interface{}) {
		if len(v) > 0 {

			for _, value := range v {
				l.send(l.ctx, l.buildlog(WARN, "", err, nil, t, value))
			}
		} else {
			l.send(l.ctx, l.buildlog(WARN, "", nil, nil, t, "empty log value passed"))
		}
	}(v...)
}

// Warnf creates a warning log using the error passed in, along with the string
// formatting and values
func (l *alog) Warnf(err error, format string, v ...interface{}) {
	t := time.Now()
	go l.send(l.ctx, l.buildlog(WARN, "", err, &format, t, v...))
}

// Errorc creates error logs based on the data coming from the
// concurrency channel that is passed in for processing
func (l *alog) Errorc(ctx context.Context, v <-chan interface{}) {
	l.clog(ctx, v, ERROR, "")
}

// Error creates an error log using the error and other values passed in
func (l *alog) Error(err error, v ...interface{}) {
	t := time.Now()
	go l.send(l.ctx, l.buildlog(ERROR, "", err, nil, t, v...))
}

// Errorln creates error logs using the error and other values passed in.
// Each error and value is printed on a different line
func (l *alog) Errorln(err error, v ...interface{}) {
	t := time.Now()
	go func(v ...interface{}) {
		if len(v) > 0 {

			for _, value := range v {
				l.send(l.ctx, l.buildlog(ERROR, "", err, nil, t, value))
			}
		} else {
			l.send(l.ctx, l.buildlog(WARN, "", nil, nil, t, "empty log value passed"))
		}
	}(v...)
}

// Errorf creates an error log using the error passed in, along with the string
// formatting and values
func (l *alog) Errorf(err error, format string, v ...interface{}) {
	t := time.Now()
	go l.send(l.ctx, l.buildlog(ERROR, "", err, &format, t, v...))
}

// Critc creates critical logs based on the data coming from the
// concurrency channel that is passed in for processing
func (l *alog) Critc(ctx context.Context, v <-chan interface{}) {
	l.clog(ctx, v, CRIT, "")
}

// Crit creates critical logs using the error and other values passed in
func (l *alog) Crit(err error, v ...interface{}) {
	t := time.Now()
	go l.send(l.ctx, l.buildlog(CRIT, "", err, nil, t, v...))
}

// Critln creates critical logs using the error and other values passed in.
// Each error and value is printed on a different line
func (l *alog) Critln(err error, v ...interface{}) {
	t := time.Now()
	go func(v ...interface{}) {
		if len(v) > 0 {

			for _, value := range v {
				l.send(l.ctx, l.buildlog(CRIT, "", err, nil, t, value))
			}
		} else {
			l.send(l.ctx, l.buildlog(WARN, "", nil, nil, t, "empty log value passed"))
		}
	}(v...)
}

// Critf creates a critical log using the error passed in, along with the string
// formatting and values
func (l *alog) Critf(err error, format string, v ...interface{}) {
	t := time.Now()
	go l.send(l.ctx, l.buildlog(CRIT, "", err, &format, t, v...))
}

// Fatalc creates fatal logs based on the data coming from the
// concurrency channel that is passed in for processing
func (l *alog) Fatalc(ctx context.Context, v <-chan interface{}) {
	l.clog(ctx, v, FATAL, "")
}

// Fatal creates a fatal log using the error and values passed into the method
func (l *alog) Fatal(err error, v ...interface{}) {
	t := time.Now()
	go l.send(l.ctx, l.buildlog(FATAL, "", err, nil, t, v...))
}

// Fatalln creates fatal logs using the error and other values passed in.
// Each error and value is printed on a different line
func (l *alog) Fatalln(err error, v ...interface{}) {
	t := time.Now()
	go func(v ...interface{}) {
		if len(v) > 0 {

			for _, value := range v {
				l.send(l.ctx, l.buildlog(FATAL, "", err, nil, t, value))
			}
		} else {
			l.send(l.ctx, l.buildlog(WARN, "", nil, nil, t, "empty log value passed"))
		}
	}(v...)
}

// Fatalf creates an error log using the error passed in, along with the string
// formatting and values
func (l *alog) Fatalf(err error, format string, v ...interface{}) {
	t := time.Now()
	go l.send(l.ctx, l.buildlog(FATAL, "", err, &format, t, v...))
}

// Customc creates custom logs based on the data coming from the
// concurrency channel that is passed in for processing
func (l *alog) Customc(ctx context.Context, v <-chan interface{}, ltype string) {
	l.clog(ctx, v, CUSTOM, ltype)
}

// Custom creates a custom log using the error and values passed into the method
func (l *alog) Custom(ltype string, err error, v ...interface{}) {
	t := time.Now()
	go l.send(l.ctx, l.buildlog(CUSTOM, ltype, err, nil, t, v...))
}

// Customln creates custom logs using the error and other values passed in.
// Each error and value is printed on a different line
func (l *alog) Customln(ltype string, err error, v ...interface{}) {
	t := time.Now()
	go func(v ...interface{}) {
		if len(v) > 0 {

			for _, value := range v {
				l.send(l.ctx, l.buildlog(CUSTOM, ltype, err, nil, t, value))
			}
		} else {
			l.send(l.ctx, l.buildlog(WARN, "", nil, nil, t, "empty log value passed"))
		}
	}(v...)
}

// Customf creates a custom log using the error passed in, along with the string
// formatting and values
func (l *alog) Customf(ltype string, err error, format string, v ...interface{}) {
	t := time.Now()
	go l.send(l.ctx, l.buildlog(CUSTOM, ltype, err, &format, t, v...))
}

// Close cancels the context of the logger internally and breaks out of
// any logging activity. This should always be called in a defer at the top
// level where the logger is initialized to ensure proper closure
func (l *alog) Close() {
	if validator.Valid(l) {

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

// Wait blocks on the logger context until the context is closed, if the close flag
// is passed then the wait function will close the context of the logger
func (l *alog) Wait(close bool) {

	// Cancel the context if indicated in the call
	if close {
		l.Close()
	}

	// Wait for all of the channels to be closed out
	<-l.cleaned
}
