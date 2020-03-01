package alog

import (
	"encoding/json"
	"strings"

	"github.com/pkg/errors"
)

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

var prefixlogsJSON = []fakelogJSON{
	{
		INFO,
		"INFO",
		nil,
		newjsonlog(
			"PREFIX",
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
			"PREFIX",
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
			"PREFIX",
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
			"PREFIX",
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
			"PREFIX",
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
			"PREFIX",
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
			"PREFIX",
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
			"PREFIX",
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
