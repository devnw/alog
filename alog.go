package alog

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"sync"
	"time"

	"github.com/benjivesterby/validator"
	"github.com/pkg/errors"
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
	buffer     int

	mutty sync.Mutex

	// The channels which will have logs sent and received on
	infodests   []chan<- log
	debugdests  []chan<- log
	warndests   []chan<- log
	errdests    []chan<- log
	critdests   []chan<- log
	fataldests  []chan<- log
	customdests []chan<- log

	// Indicates that all logs have been cleared to the respective
	// destinations during a close
	closed chan bool
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

	go func(ctx context.Context, logs <-chan log, destination Destination) {
		// TODO: handle panic

		for {
			select {
			case <-ctx.Done():
				// TODO: setup to close the destination if it has a close method
				return
			case l, ok := <-logs:
				if ok {
					if validator.IsValid(l) {
						var message string

						switch destination.Format {
						case DELIM:
						case JSON:
							if msg, err := json.Marshal(l); err == nil {
								message = string(msg)
							} else {
								// TODO: panic?
							}
						default:
							message = l.String()
						}

						fmt.Fprint(destination.Writer, string(message))
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
func (l *alog) send(ctx context.Context, value log) (err error) {
	// TODO: Handle panic here

	// Break out in the event that the context has been cancelled
	select {
	case <-ctx.Done():
		err = errors.New("logger is closed")
	default:
		// Loop over the destinations for this logtype and push onto the
		// log channels for each destination
		for _, destination := range l.getd(value.logtype) {

			// Push the log onto the destination channel
			select {
			case <-ctx.Done():
				destination <- value
				close(destination)
			case destination <- value:
			}
		}
	}

	return err
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

func (l *alog) clog(ctx context.Context, v <-chan interface{}, logtype int8, custom string) {

	go func(ctx context.Context, v <-chan interface{}, logtype int8, custom string) {
		// TODO: handle panic
		for {
			select {
			case <-ctx.Done():
				return
			case value, ok := <-v:
				if ok {
					switch t := value.(type) {
					case error:
						l.send(ctx, l.buildlog(logtype, custom, t, nil, nil))
					default:
						l.send(ctx, l.buildlog(logtype, custom, nil, nil, []interface{}{t}))
					}
				} else {
					return
				}
			}
		}
	}(ctx, v, logtype, custom)
}

// Printc creates informational logs based on the data coming from the
// concurrency channel that is passed in for processing
func (l *alog) Printc(ctx context.Context, v <-chan interface{}) {
	l.clog(ctx, v, INFO, "")
}

// Print creates informational logs based on the inputs
func (l *alog) Print(v ...interface{}) error {
	return l.send(l.ctx, l.buildlog(INFO, "", nil, nil, v...))
}

// Println prints the data coming in as an informational log on individual lines
func (l *alog) Println(v ...interface{}) (out error) {
	for _, value := range v {
		if out = l.send(l.ctx, l.buildlog(INFO, "", nil, nil, value)); out != nil {

			// Break the loop and return the error
			break
		}
	}

	return out
}

// Printf creates an informational log using the format and values
func (l *alog) Printf(format string, v ...interface{}) error {
	return l.send(l.ctx, l.buildlog(INFO, "", nil, &format, v...))
}

// Debugc creates debug logs based on the data coming from the
// concurrency channel that is passed in for processing
func (l *alog) Debugc(ctx context.Context, v <-chan interface{}) {
	l.clog(ctx, v, DEBUG, "")
}

// Debug creates debugging logs based on the inputs
func (l *alog) Debug(err error, v ...interface{}) error {
	return l.send(l.ctx, l.buildlog(DEBUG, "", err, nil, v...))
}

// Debugln prints the data coming in as a debug log on individual lines
func (l *alog) Debugln(err error, v ...interface{}) (out error) {
	for _, value := range v {
		if out = l.send(l.ctx, l.buildlog(DEBUG, "", err, nil, value)); out != nil {

			// Break the loop and return the error
			break
		}
	}

	return out
}

// Debugf creates an debugging log using the format and values
func (l *alog) Debugf(err error, format string, v ...interface{}) error {
	return l.send(l.ctx, l.buildlog(DEBUG, "", err, &format, v...))
}

// Warnc creates warning logs based on the data coming from the
// concurrency channel that is passed in for processing
func (l *alog) Warnc(ctx context.Context, v <-chan interface{}) {
	l.clog(ctx, v, WARN, "")
}

// Warn creates a warning log using the error passed in along with the
// values passed in
func (l *alog) Warn(err error, v ...interface{}) error {
	return l.send(l.ctx, l.buildlog(WARN, "", err, nil, v...))
}

// Warnln creates a warning log using the error and values passed in.
// Each error and value is printed on a different line
func (l *alog) Warnln(err error, v ...interface{}) (out error) {
	for _, value := range v {
		if out = l.send(l.ctx, l.buildlog(WARN, "", err, nil, value)); out != nil {

			// Break the loop and return the error
			break
		}
	}

	return out
}

// Warnf creates a warning log using the error passed in, along with the string
// formatting and values
func (l *alog) Warnf(err error, format string, v ...interface{}) error {
	return l.send(l.ctx, l.buildlog(WARN, "", err, &format, v...))
}

// Errorc creates error logs based on the data coming from the
// concurrency channel that is passed in for processing
func (l *alog) Errorc(ctx context.Context, v <-chan interface{}) {
	l.clog(ctx, v, ERROR, "")
}

// Error creates an error log using the error and other values passed in
func (l *alog) Error(err error, v ...interface{}) error {
	return l.send(l.ctx, l.buildlog(ERROR, "", err, nil, v...))
}

// Errorln creates error logs using the error and other values passed in.
// Each error and value is printed on a different line
func (l *alog) Errorln(err error, v ...interface{}) (out error) {
	for _, value := range v {
		if out = l.send(l.ctx, l.buildlog(ERROR, "", err, nil, value)); out != nil {

			// Break the loop and return the error
			break
		}
	}

	return out
}

// Errorf creates an error log using the error passed in, along with the string
// formatting and values
func (l *alog) Errorf(err error, format string, v ...interface{}) error {
	return l.send(l.ctx, l.buildlog(ERROR, "", err, &format, v...))
}

// Critc creates critical logs based on the data coming from the
// concurrency channel that is passed in for processing
func (l *alog) Critc(ctx context.Context, v <-chan interface{}) {
	l.clog(ctx, v, CRIT, "")
}

// Crit creates critical logs using the error and other values passed in
func (l *alog) Crit(err error, v ...interface{}) error {
	return l.send(l.ctx, l.buildlog(CRIT, "", err, nil, v...))
}

// Critln creates critical logs using the error and other values passed in.
// Each error and value is printed on a different line
func (l *alog) Critln(err error, v ...interface{}) (out error) {
	for _, value := range v {
		if out = l.send(l.ctx, l.buildlog(CRIT, "", err, nil, value)); out != nil {

			// Break the loop and return the error
			break
		}
	}

	return out
}

// Critf creates a critical log using the error passed in, along with the string
// formatting and values
func (l *alog) Critf(err error, format string, v ...interface{}) error {
	return l.send(l.ctx, l.buildlog(CRIT, "", err, &format, v...))
}

func (l *alog) fpanic(err error, out error) {

	if out != nil && err != nil {
		err = errors.Errorf("error while sending fatal log %s | log: %s", out.Error(), err.Error())
	} else if err == nil {
		err = errors.New("fatal error panic")
	}

	// TODO: Update panic to include information about the fatal, as well as stack trace information
	panic(err)
}

// Fatal creates a fatal log using the error and values passed into the method
// After logging the fatal log the Fatal method throws a panic to crash the application
func (l *alog) Fatal(err error, v ...interface{}) {
	defer l.cancel()
	out := l.send(l.ctx, l.buildlog(FATAL, "", err, nil, v...))

	l.fpanic(err, out)
}

// Fatalln creates fatal logs using the error and other values passed in.
// Each error and value is printed on a different line
// After logging the fatal log the Fatalln method throws a panic to crash the application
func (l *alog) Fatalln(err error, v ...interface{}) {
	defer l.cancel()

	var out error
	for _, value := range v {
		if out = l.send(l.ctx, l.buildlog(FATAL, "", err, nil, value)); out != nil {

			// Break the loop and return the error
			break
		}
	}

	l.fpanic(err, out)
}

// Fatalf creates an error log using the error passed in, along with the string
// formatting and values
// After logging the fatal log the Fatalf method throws a panic to crash the application
func (l *alog) Fatalf(err error, format string, v ...interface{}) {
	defer l.cancel()

	out := l.send(l.ctx, l.buildlog(FATAL, "", err, &format, v...))

	l.fpanic(err, out)
}

// Customc creates custom logs based on the data coming from the
// concurrency channel that is passed in for processing
func (l *alog) Customc(ctx context.Context, v <-chan interface{}, ltype string) {
	l.clog(ctx, v, CUSTOM, ltype)
}

// Custom creates a custom log using the error and values passed into the method
func (l *alog) Custom(ltype string, err error, v ...interface{}) error {
	return l.send(l.ctx, l.buildlog(CUSTOM, ltype, err, nil, v...))
}

// Customln creates custom logs using the error and other values passed in.
// Each error and value is printed on a different line
func (l *alog) Customln(ltype string, err error, v ...interface{}) (out error) {
	for _, value := range v {
		if out = l.send(l.ctx, l.buildlog(CUSTOM, ltype, err, nil, value)); out != nil {

			// Break the loop and return the error
			break
		}
	}

	return out
}

// Customf creates a custom log using the error passed in, along with the string
// formatting and values
func (l *alog) Customf(ltype string, err error, format string, v ...interface{}) error {
	return l.send(l.ctx, l.buildlog(CUSTOM, ltype, err, &format, v...))
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

// Wait blocks on the logger context until the context is closed
func (l *alog) Wait() {

	// Cancel the internal context
	l.cancel()

	// TODO: Wait for all of the channels to be closed out
	<-l.ctx.Done()
}
