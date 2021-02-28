package alog

import (
	"context"
	"errors"
	"fmt"
	"io"
	"strings"
	"time"
)

type passmock struct {
	msg chan []byte
}

func check(value []byte, expected string) (err error) {
	if len(value) <= 0 {
		return fmt.Errorf("value is empty")
	}

	if value[len(value)-1] != '\n' {
		return fmt.Errorf("expected newline at end of log")
	}

	output := string(value)
	if strings.Contains(output, "PREFIX") {
		begin := 7 // length of prefix
		end := strings.Index(output, "[")
		output = strings.TrimSpace(output[:begin] + output[end:])
	} else {
		i := strings.Index(output, "[")
		output = strings.TrimSpace(output[i:])
	}

	if expected != output {
		return fmt.Errorf("expected result: '%s' != output: '%s'", expected, output)
	}

	return nil
}

var lvls = INFO | DEBUG | TRACE | WARN | ERROR | CRIT | FATAL | CUSTOM

func (pm *passmock) Write(p []byte) (n int, err error) {

	n = len(p)
	pm.msg <- p

	return n, err
}

func testg(dest *Destination, w io.Writer) error {

	if dest == nil {

		if w == nil {
			return errors.New("nil io.Writer passsed to testg")
		}

		dest = &Destination{
			lvls,
			STD,
			w,
		}
	}
	return Global(
		context.Background(),
		"",
		DEFAULTTIMEFORMAT,
		time.UTC,
		DEFAULTBUFFER,
		*dest,
	)
}

var std = map[LogLevel]func(err error, v ...interface{}){
	DEBUG: Debug,
	TRACE: Trace,
	WARN:  Warn,
	ERROR: Error,
	CRIT:  Crit,
	FATAL: Fatal,
}

var stdlns = map[LogLevel]func(err error, v ...interface{}){
	DEBUG: Debugln,
	TRACE: Traceln,
	WARN:  Warnln,
	ERROR: Errorln,
	CRIT:  Critln,
	FATAL: Fatalln,
}

var stdfs = map[LogLevel]func(err error, format string, v ...interface{}){
	DEBUG: Debugf,
	TRACE: Tracef,
	WARN:  Warnf,
	ERROR: Errorf,
	CRIT:  Critf,
	FATAL: Fatalf,
}

var chanfs = map[LogLevel]func(ctx context.Context, v <-chan interface{}){
	INFO:  Printc,
	DEBUG: Debugc,
	TRACE: Tracec,
	WARN:  Warnc,
	ERROR: Errorc,
	CRIT:  Critc,
	FATAL: Fatalc,
}

type fakelog struct {
	lvl      LogLevel
	text     string
	err      error
	expected string
}

type abnorm struct {
	value string
}

type iny struct {
	value interface{}
}

var logs = []fakelog{
	{
		INFO,
		"INFO",
		nil,
		"[INFO] INFO",
	},
	{
		DEBUG,
		"DEBUG",
		errors.New("DEBUG"),
		"[DEBUG] DEBUG | err: DEBUG",
	},
	{
		TRACE,
		"TRACE",
		errors.New("TRACE"),
		"[TRACE] TRACE | err: TRACE",
	},
	{
		WARN,
		"WARN",
		errors.New("WARN"),
		"[WARN] WARN | err: WARN",
	},
	{
		ERROR,
		"ERROR",
		errors.New("ERROR"),
		"[ERROR] ERROR | err: ERROR",
	},
	{
		CRIT,
		"CRIT",
		errors.New("CRIT"),
		"[CRITICAL] CRIT | err: CRIT",
	},
	{
		FATAL,
		"FATAL",
		errors.New("FATAL"),
		"[FATAL] FATAL | err: FATAL",
	},
	{
		CUSTOM,
		"CUSTOM",
		errors.New("CUSTOM"),
		"[CUSTOM] CUSTOM | err: CUSTOM",
	},
}

var prefixlogs = []fakelog{
	{
		INFO,
		"INFO",
		nil,
		"PREFIX [INFO] INFO",
	},
	{
		DEBUG,
		"DEBUG",
		errors.New("DEBUG"),
		"PREFIX [DEBUG] DEBUG | err: DEBUG",
	},
	{
		TRACE,
		"TRACE",
		errors.New("TRACE"),
		"PREFIX [TRACE] TRACE | err: TRACE",
	},
	{
		WARN,
		"WARN",
		errors.New("WARN"),
		"PREFIX [WARN] WARN | err: WARN",
	},
	{
		ERROR,
		"ERROR",
		errors.New("ERROR"),
		"PREFIX [ERROR] ERROR | err: ERROR",
	},
	{
		CRIT,
		"CRIT",
		errors.New("CRIT"),
		"PREFIX [CRITICAL] CRIT | err: CRIT",
	},
	{
		FATAL,
		"FATAL",
		errors.New("FATAL"),
		"PREFIX [FATAL] FATAL | err: FATAL",
	},
	{
		CUSTOM,
		"CUSTOM",
		errors.New("CUSTOM"),
		"PREFIX [CUSTOM] CUSTOM | err: CUSTOM",
	},
}

var cerrlogs = []fakelog{
	{
		DEBUG,
		"DEBUG",
		errors.New("DEBUG"),
		"[DEBUG] stream log | err: DEBUG",
	},
	{
		TRACE,
		"TRACE",
		errors.New("TRACE"),
		"[TRACE] stream log | err: TRACE",
	},
	{
		WARN,
		"WARN",
		errors.New("WARN"),
		"[WARN] stream log | err: WARN",
	},
	{
		ERROR,
		"ERROR",
		errors.New("ERROR"),
		"[ERROR] stream log | err: ERROR",
	},
	{
		CRIT,
		"CRIT",
		errors.New("CRIT"),
		"[CRITICAL] stream log | err: CRIT",
	},
	{
		FATAL,
		"FATAL",
		errors.New("FATAL"),
		"[FATAL] stream log | err: FATAL",
	},
	{
		CUSTOM,
		"CUSTOM",
		errors.New("CUSTOM"),
		"[CUSTOM] stream log | err: CUSTOM",
	},
}

var clogs = []fakelog{
	{
		INFO,
		"INFO",
		nil,
		"[INFO] INFO",
	},
	{
		DEBUG,
		"DEBUG",
		errors.New("DEBUG"),
		"[DEBUG] DEBUG",
	},
	{
		TRACE,
		"TRACE",
		errors.New("TRACE"),
		"[TRACE] TRACE",
	},
	{
		WARN,
		"WARN",
		errors.New("WARN"),
		"[WARN] WARN",
	},
	{
		ERROR,
		"ERROR",
		errors.New("ERROR"),
		"[ERROR] ERROR",
	},
	{
		CRIT,
		"CRIT",
		errors.New("CRIT"),
		"[CRITICAL] CRIT",
	},
	{
		FATAL,
		"FATAL",
		errors.New("FATAL"),
		"[FATAL] FATAL",
	},
	{
		CUSTOM,
		"CUSTOM",
		errors.New("CUSTOM"),
		"[CUSTOM] CUSTOM",
	},
}

var multi = []fakelog{
	{
		INFO,
		"INFO",
		nil,
		"[INFO] INFO,INFO",
	},
	{
		DEBUG,
		"DEBUG",
		errors.New("DEBUG"),
		"[DEBUG] DEBUG,DEBUG | err: DEBUG",
	},
	{
		TRACE,
		"TRACE",
		errors.New("TRACE"),
		"[TRACE] TRACE,TRACE | err: TRACE",
	},
	{
		WARN,
		"WARN",
		errors.New("WARN"),
		"[WARN] WARN,WARN | err: WARN",
	},
	{
		ERROR,
		"ERROR",
		errors.New("ERROR"),
		"[ERROR] ERROR,ERROR | err: ERROR",
	},
	{
		CRIT,
		"CRIT",
		errors.New("CRIT"),
		"[CRITICAL] CRIT,CRIT | err: CRIT",
	},
	{
		FATAL,
		"FATAL",
		errors.New("FATAL"),
		"[FATAL] FATAL,FATAL | err: FATAL",
	},
	{
		CUSTOM,
		"CUSTOM",
		errors.New("CUSTOM"),
		"[CUSTOM] CUSTOM,CUSTOM | err: CUSTOM",
	},
}

var flogs = []fakelog{
	{
		INFO,
		"INFO",
		nil,
		"[INFO] INFO *INFO*",
	},
	{
		DEBUG,
		"DEBUG",
		errors.New("DEBUG"),
		"[DEBUG] DEBUG *DEBUG* | err: DEBUG",
	},
	{
		TRACE,
		"TRACE",
		errors.New("TRACE"),
		"[TRACE] TRACE *TRACE* | err: TRACE",
	},
	{
		WARN,
		"WARN",
		errors.New("WARN"),
		"[WARN] WARN *WARN* | err: WARN",
	},
	{
		ERROR,
		"ERROR",
		errors.New("ERROR"),
		"[ERROR] ERROR *ERROR* | err: ERROR",
	},
	{
		CRIT,
		"CRIT",
		errors.New("CRIT"),
		"[CRITICAL] CRIT *CRIT* | err: CRIT",
	},
	{
		FATAL,
		"FATAL",
		errors.New("FATAL"),
		"[FATAL] FATAL *FATAL* | err: FATAL",
	},
	{
		CUSTOM,
		"CUSTOM",
		errors.New("CUSTOM"),
		"[CUSTOM] CUSTOM *CUSTOM* | err: CUSTOM",
	},
}
