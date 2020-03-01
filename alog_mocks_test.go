package alog

import (
	"context"
	"encoding/json"
	"io"
	"strings"
	"time"

	"github.com/pkg/errors"
)

type passmock struct {
	msg chan []byte
}

func check(value []byte, expected string) (err error) {
	if len(value) > 0 {

		output := string(value)

		if strings.LastIndex(output, "\n") == len(output)-1 {
			i := strings.Index(output, "[")
			output = strings.TrimSpace(output[i:])

			if expected != output {
				err = errors.Errorf("expected result: '%s' != output: '%s'", expected, output)
			}

		} else {
			err = errors.Errorf("expected newline at end of log")
		}
	} else {
		err = errors.Errorf("value is empty")
	}

	return err
}

func checkJSON(value []byte, expected jsonlogtest) (err error) {

	if len(value) > 0 {

		output := string(value)

		if strings.LastIndex(output, "\n") == len(output)-1 {

			received := jsonlogtest{}
			if err = json.Unmarshal(value, &received); err == nil {
				if !same(expected, received) {
					err = errors.Errorf("expected: '%s' does not match received: '%s'", expected.String(), string(value))
				}
			}
		} else {
			err = errors.Errorf("expected newline at end of log")
		}
	} else {
		err = errors.Errorf("value is empty")
	}

	return err
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

type fakelogJSON struct {
	lvl      LogLevel
	text     string
	err      error
	expected jsonlogtest
}

func same(obj1, obj2 jsonlogtest) (same bool) {

	// Prefix check
	if obj1.Prefix == nil &&
		obj2.Prefix == nil ||
		(obj1.Prefix != nil &&
			obj2.Prefix != nil &&
			strings.Compare(*obj1.Prefix, *obj2.Prefix) == 0) {

		if strings.Compare(obj1.LogType, obj2.LogType) == 0 {

			if obj1.Error == nil &&
				obj2.Error == nil ||
				(obj1.Error != nil &&
					obj2.Error != nil &&
					strings.Compare(*obj1.Error, *obj2.Error) == 0) {

				if len(obj1.Messages) == len(obj2.Messages) {
					for i := range obj1.Messages {
						if strings.Compare(obj1.Messages[i], obj2.Messages[i]) != 0 {
							return
						}
					}

					same = true
				}
			}
		}
	}

	return same
}

type jsonlogtest struct {
	Prefix    *string  `json:"prefix,omitempty"`
	LogType   string   `json:"type"`
	Timestamp string   `json:"timestamp"`
	Error     *string  `json:"error,omitempty"`
	Messages  []string `json:"messages"`
}

func (j jsonlogtest) String() string {
	if out, err := json.Marshal(&j); err != nil {
		panic("error marshalling jsonlogtest")
	} else {
		return string(out)
	}
}

func newjsonlog(pfix string, ltype string, err string, msgs ...string) jsonlogtest {

	var p *string
	if len(pfix) > 0 {
		p = &pfix
	}

	var e *string
	if len(err) > 0 {
		e = &err
	}

	return jsonlogtest{
		p,
		ltype,
		"",
		e,
		msgs,
	}
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

var logsJSON = []fakelogJSON{
	{
		INFO,
		"INFO",
		nil,
		newjsonlog(
			"",
			"INFO",
			"",
			"INFO",
		),
	},
	{
		DEBUG,
		"DEBUG",
		errors.New("DEBUG"),
		newjsonlog(
			"",
			"DEBUG",
			"DEBUG",
			"DEBUG",
		),
	},
	{
		TRACE,
		"TRACE",
		errors.New("TRACE"),
		newjsonlog(
			"",
			"TRACE",
			"TRACE",
			"TRACE",
		),
	},
	{
		WARN,
		"WARN",
		errors.New("WARN"),
		newjsonlog(
			"",
			"WARN",
			"WARN",
			"WARN",
		),
	},
	{
		ERROR,
		"ERROR",
		errors.New("ERROR"),
		newjsonlog(
			"",
			"ERROR",
			"ERROR",
			"ERROR",
		),
	},
	{
		CRIT,
		"CRIT",
		errors.New("CRIT"),
		newjsonlog(
			"",
			"CRITICAL",
			"CRIT",
			"CRIT",
		),
	},
	{
		FATAL,
		"FATAL",
		errors.New("FATAL"),
		newjsonlog(
			"",
			"FATAL",
			"FATAL",
			"FATAL",
		),
	},
	{
		CUSTOM,
		"CUSTOM",
		errors.New("CUSTOM"),
		newjsonlog(
			"",
			"CUSTOM",
			"CUSTOM",
			"CUSTOM",
		),
	},
}

var cerrlogsJSON = []fakelogJSON{
	{
		DEBUG,
		"DEBUG",
		errors.New("DEBUG"),
		newjsonlog(
			"",
			"DEBUG",
			"DEBUG",
			"stream log",
		),
	},
	{
		TRACE,
		"TRACE",
		errors.New("TRACE"),
		newjsonlog(
			"",
			"TRACE",
			"TRACE",
			"stream log",
		),
	},
	{
		WARN,
		"WARN",
		errors.New("WARN"),
		newjsonlog(
			"",
			"WARN",
			"WARN",
			"stream log",
		),
	},
	{
		ERROR,
		"ERROR",
		errors.New("ERROR"),
		newjsonlog(
			"",
			"ERROR",
			"ERROR",
			"stream log",
		),
	},
	{
		CRIT,
		"CRIT",
		errors.New("CRIT"),
		newjsonlog(
			"",
			"CRITICAL",
			"CRIT",
			"stream log",
		),
	},
	{
		FATAL,
		"FATAL",
		errors.New("FATAL"),
		newjsonlog(
			"",
			"FATAL",
			"FATAL",
			"stream log",
		),
	},
	{
		CUSTOM,
		"CUSTOM",
		errors.New("CUSTOM"),
		newjsonlog(
			"",
			"CUSTOM",
			"CUSTOM",
			"stream log",
		),
	},
}

var clogsJSON = []fakelogJSON{
	{
		INFO,
		"INFO",
		nil,
		newjsonlog(
			"",
			"INFO",
			"",
			"INFO",
		),
	},
	{
		DEBUG,
		"DEBUG",
		errors.New("DEBUG"),
		newjsonlog(
			"",
			"DEBUG",
			"",
			"DEBUG",
		),
	},
	{
		TRACE,
		"TRACE",
		errors.New("TRACE"),
		newjsonlog(
			"",
			"TRACE",
			"",
			"TRACE",
		),
	},
	{
		WARN,
		"WARN",
		errors.New("WARN"),
		newjsonlog(
			"",
			"WARN",
			"",
			"WARN",
		),
	},
	{
		ERROR,
		"ERROR",
		errors.New("ERROR"),
		newjsonlog(
			"",
			"ERROR",
			"",
			"ERROR",
		),
	},
	{
		CRIT,
		"CRIT",
		errors.New("CRIT"),
		newjsonlog(
			"",
			"CRITICAL",
			"",
			"CRIT",
		),
	},
	{
		FATAL,
		"FATAL",
		errors.New("FATAL"),
		newjsonlog(
			"",
			"FATAL",
			"",
			"FATAL",
		),
	},
	{
		CUSTOM,
		"CUSTOM",
		errors.New("CUSTOM"),
		newjsonlog(
			"",
			"CUSTOM",
			"",
			"CUSTOM",
		),
	},
}

var multiJSON = []fakelogJSON{
	{
		INFO,
		"INFO",
		nil,
		newjsonlog(
			"",
			"INFO",
			"",
			"INFO",
			"INFO",
		),
	},
	{
		DEBUG,
		"DEBUG",
		errors.New("DEBUG"),
		newjsonlog(
			"",
			"DEBUG",
			"DEBUG",
			"DEBUG",
			"DEBUG",
		),
	},
	{
		TRACE,
		"TRACE",
		errors.New("TRACE"),
		newjsonlog(
			"",
			"TRACE",
			"TRACE",
			"TRACE",
			"TRACE",
		),
	},
	{
		WARN,
		"WARN",
		errors.New("WARN"),
		newjsonlog(
			"",
			"WARN",
			"WARN",
			"WARN",
			"WARN",
		),
	},
	{
		ERROR,
		"ERROR",
		errors.New("ERROR"),
		newjsonlog(
			"",
			"ERROR",
			"ERROR",
			"ERROR",
			"ERROR",
		),
	},
	{
		CRIT,
		"CRIT",
		errors.New("CRIT"),
		newjsonlog(
			"",
			"CRITICAL",
			"CRIT",
			"CRIT",
			"CRIT",
		),
	},
	{
		FATAL,
		"FATAL",
		errors.New("FATAL"),
		newjsonlog(
			"",
			"FATAL",
			"FATAL",
			"FATAL",
			"FATAL",
		),
	},
	{
		CUSTOM,
		"CUSTOM",
		errors.New("CUSTOM"),
		newjsonlog(
			"",
			"CUSTOM",
			"CUSTOM",
			"CUSTOM",
			"CUSTOM",
		),
	},
}

var flogsJSON = []fakelogJSON{
	{
		INFO,
		"INFO",
		nil,
		newjsonlog(
			"",
			"INFO",
			"",
			"INFO *INFO*",
		),
	},
	{
		DEBUG,
		"DEBUG",
		errors.New("DEBUG"),
		newjsonlog(
			"",
			"DEBUG",
			"DEBUG",
			"DEBUG *DEBUG*",
		),
	},
	{
		TRACE,
		"TRACE",
		errors.New("TRACE"),
		newjsonlog(
			"",
			"TRACE",
			"TRACE",
			"TRACE *TRACE*",
		),
	},
	{
		WARN,
		"WARN",
		errors.New("WARN"),
		newjsonlog(
			"",
			"WARN",
			"WARN",
			"WARN *WARN*",
		),
	},
	{
		ERROR,
		"ERROR",
		errors.New("ERROR"),
		newjsonlog(
			"",
			"ERROR",
			"ERROR",
			"ERROR *ERROR*",
		),
	},
	{
		CRIT,
		"CRIT",
		errors.New("CRIT"),
		newjsonlog(
			"",
			"CRITICAL",
			"CRIT",
			"CRIT *CRIT*",
		),
	},
	{
		FATAL,
		"FATAL",
		errors.New("FATAL"),
		newjsonlog(
			"",
			"FATAL",
			"FATAL",
			"FATAL *FATAL*",
		),
	},
	{
		CUSTOM,
		"CUSTOM",
		errors.New("CUSTOM"),
		newjsonlog(
			"",
			"CUSTOM",
			"CUSTOM",
			"CUSTOM *CUSTOM*",
		),
	},
}
