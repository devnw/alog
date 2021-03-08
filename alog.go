// Copyright © 2019 Developer Network, LLC
//
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of this source code package.

package alog

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"devnw.com/validator"
)

// TODO: Setup so that new destinations can be added at runtime (chan Dest)
// TODO: flag for stack traces on logs with errors?

// streamlog is a constant for the log value of a streaming log when
// an error type is sent
const streamlog = "stream log"

type alog struct {
	ctx          context.Context
	cancel       context.CancelFunc
	destinations []Destination

	// location is the timezone that is used for logging.
	// Default: UTC
	location *time.Location

	// dateformat is the date format that is used in the logging.
	// Default: RFC3339
	dateformat string
	prefix     string
	buffer     int

	mutty sync.RWMutex

	out map[LogLevel][]chan<- *log

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

	// Loop over the destinations and close the channels
	for _, destinations := range l.out {
		for _, out := range destinations {
			ctx, cancel := context.WithTimeout(
				l.ctx,
				time.Millisecond,
			)

			// Wait for the channel to empty or timeout
			// to elapse, whichever is sooner
			go func(
				ctx context.Context,
				cancel context.CancelFunc,
				out chan<- *log,
			) {
				defer close(out)
				defer cancel()

				for len(out) > 0 {
					select {
					case <-ctx.Done():
					default:
					}
				}
			}(ctx, cancel, out)
		}
	}
}

// init starts up the go routines for receiving and publishing logs
// to the available io.Writers
func (l *alog) init() {
	// Startup the cleanup go routine to monitor
	// for the closed context switch
	go l.cleanup()

	for _, dest := range l.destinations {
		for level := range l.out {
			if dest.Types&level <= 0 {
				continue
			}

			l.out[level] = append(
				l.out[level],
				l.listen(l.ctx, dest),
			)
		}
	}
}

func (l *alog) listen(
	ctx context.Context,
	destination Destination,
) chan<- *log {
	logs := make(chan *log)

	go func(ctx context.Context, logs <-chan *log, destination Destination) {
		// TODO: handle panic

		for {
			select {
			case <-ctx.Done():
				// TODO: setup to close the destination
				// if it has a close method
				return
			case l, ok := <-logs:
				if !ok {
					return
				}

				if !validator.Valid(l) {
					continue
				}

				message := l.String()
				if destination.Format == JSON {
					msg, err := json.Marshal(l)
					if err != nil {
						// TODO: panic?
						panic("error marshaling JSON")
					}

					// Add a newline to each json log for
					// readability
					message = string(msg) + "\n"
				}

				_, err := destination.Writer.Write([]byte(message))
				if err != nil {
					panic("error writing to destination")
				}
			}
		}
	}(ctx, logs, destination)

	return logs
}

// send is used to create a go routine thread for fanning out specific
// log types to each of the destinations
func (l *alog) send(ctx context.Context, value *log) {
	// TODO: Handle panic here
	if value == nil {
		return
	}

	// Break out in the event that the context has been canceled
	select {
	case <-ctx.Done():
	default:
		// Lock reads here while pulling channels
		l.mutty.RLock()
		defer l.mutty.RUnlock()

		// Loop over the destinations for this logtype and push onto the
		// log channels for each destination
		for _, destination := range l.out[value.logtype] {
			// Push the log onto the destination channel
			select {
			case <-ctx.Done():
				return
			case destination <- value:
			}
		}
	}
}

func (l *alog) buildlog(
	logtype LogLevel,
	custom string,
	err error,
	format *string,
	t time.Time,
	v ...interface{},
) (newlog *log) {
	values := v
	if format != nil {
		values = []interface{}{fmt.Sprintf(*format, v...)}
	}

	newlog = &log{
		logger:     l,
		logtype:    logtype,
		customtype: custom,
		timestamp:  t,
		err:        err,
		values:     values,
	}

	return newlog
}

func (l *alog) clog(
	ctx context.Context,
	v <-chan interface{},
	level LogLevel,
	custom string,
) {
	go func(
		ctx context.Context,
		v <-chan interface{},
		level LogLevel,
		custom string,
	) {
		// TODO: handle panic
		for {
			select {
			case <-ctx.Done():
				return
			case value, ok := <-v:
				if !ok {
					return
				}

				switch t := value.(type) {
				case nil:
				case error:
					l.send(
						ctx,
						l.buildlog(
							level,
							custom,
							t,
							nil,
							time.Now(),
							streamlog,
						),
					)
				default:
					l.send(
						ctx,
						l.buildlog(
							level,
							custom,
							nil,
							nil,
							time.Now(),
							t,
						),
					)
				}
			}
		}
	}(ctx, v, level, custom)
}

func (l *alog) sendMultiLine(
	level LogLevel,
	err error,
	values ...interface{},
) {
	t := time.Now()
	for _, value := range values {
		l.send(
			l.ctx,
			l.buildlog(level, "", err, nil, t, value),
		)
	}
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
	go l.sendMultiLine(INFO, nil, v...)
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
	go l.sendMultiLine(DEBUG, err, v...)
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
	go l.sendMultiLine(TRACE, err, v...)
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
	go l.sendMultiLine(WARN, err, v...)
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
	go l.sendMultiLine(ERROR, err, v...)
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
	go l.sendMultiLine(CRIT, err, v...)
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
	go l.sendMultiLine(FATAL, err, v...)
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
		for _, value := range v {
			l.send(
				l.ctx,
				l.buildlog(
					CUSTOM,
					ltype,
					err,
					nil,
					t,
					value,
				),
			)
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
func (l *alog) Wait(exit bool) {
	// Cancel the context if indicated in the call
	if exit {
		l.Close()
	}

	// Wait for all of the channels to be closed out
	<-l.cleaned
}
