# Fully concurrent, non-IO blocking Logger for Go

[![CI](https://github.com/devnw/alog/workflows/CI/badge.svg)](https://github.com/devnw/alog/actions)
[![Go Report Card](https://goreportcard.com/badge/devnw.com/alog)](https://goreportcard.com/report/devnw.com/alog)
[![codecov](https://codecov.io/gh/devnw/alog/branch/master/graph/badge.svg)](https://codecov.io/gh/devnw/alog)
[![Go Reference](https://pkg.go.dev/badge/devnw.com/alog.svg)](https://pkg.go.dev/devnw.com/alog)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg)](http://makeapullrequest.com)

## Usage

Installation:

```go
go get -u devnw.com/alog
```

## alog

is an implementation of a fully concurrent non-IO blocking
log library. Internally messages are routed through a number of subscribed
destinations consisting of destination log levels, timezone formatting
strings and an io.Writer for the destination to be written to.

The globally available functions utilize a singleton implementation of a
global logger that is initialized when the package is loaded. This global
logger utilizes os.Stdout for INFO, DEBUG, TRACE, WARN, CUSTOM logs, and
os.Stderr for ERROR, CRITICAL, FATAL as defined in the method Standard()

The global logger singleton can be overridden using the "Global" method
passing in the proper parameters as required by your use case.

If you prefer to use a non-global singleton logger you can use the "New"
method to create a new logger which implements the alog packages "Logger"
interface which can be passed throughout an application as a logger.

Each log level has an associated function:

* INFO - Print, Printf, Println, Printc
* DEBUG - Debug, Debugf, Debugln, Debugc
* TRACE - Trace, Tracef, Traceln, Tracec
* WARN - Warn, Warnf, Warnln, Warnc
* ERROR - Error, Errorf, Errorln, Errorc
* CRITICAL - Crit, Critf, Critln, Critc
* FATAL - Fatal, Fatalf, Fatalln, Fatalc
* CUSTOM Level - Custom, Customf, Customln, Customc

c(channel) functions are special methods in the logger that rather than being called
directly are passed a channel which receives logs into the logger in a
concurrent manner without direct calls to the `c(channel) function` again afterwards.
These methods are particularly useful for applications where the overhead of
a logger in concurrent actions is a particular nuisance. For these cases the
c(channel) functions are a perfect use case. Each log level supports it's own
c(channel) function.

This library currently support two log formatting types

* STD: Formats your logs using a "PREFIX DATE [LEVEL] Log | err: Error"
* JSON: Formats your logs using JSON in the following form

```json
{
    "prefix":"PREFIX",
    "type":"ERROR",
    "timestamp":"2020-03-01T16:21:28-06:00",
    "error":"Error Message",
    "messages":["Log Message 1", "Log Message 2"]
}
```

Global Usage:

```go

alog.Println("info log")
alog.Debugln(err, "debug log")
alog.Traceln(err, "trace log")
alog.Warnln(err, "warn log")
alog.Errorln(err, "err log")
alog.Critln(err, "critcal log")
alog.Fatalln(err, "fatal log")
alog.Customln("CUSTOM", err, "debug log")

```

For applications that want to ensure log completeness using `alog` execute
`alog.Wait(bool)` when cleaning up your application so that the logger can
correctly cleanup any lingering channels and Go routines. The boolean
parameter let's the wait method know if you want to also close the logger internally.
