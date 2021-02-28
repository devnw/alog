// Package alog is an implementation of a fully concurrent non-IO blocking
// log library. Internally messages are routed through a number of subscribed
// destinations consisting of destination log levels, timezone formatting
// strings and an io.Writer for the destination to be written to.
//
// The globally available functions utilize a singleton implementation of a
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
