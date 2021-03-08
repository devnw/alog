package alog

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func checkJSON(value []byte, expected jsonlogtest) (err error) {
	if len(value) == 0 {
		return fmt.Errorf("value is empty")
	}

	if value[len(value)-1] != '\n' {
		return fmt.Errorf("expected newline at end of log")
	}

	received := jsonlogtest{}
	err = json.Unmarshal(value, &received)
	if err != nil {
		return err
	}

	diff, ok := same(expected, received)
	if !ok {
		return fmt.Errorf("expected same, but got \n %s", diff)
	}

	return nil
}

type fakelogJSON struct {
	lvl      LogLevel
	text     string
	err      error
	expected jsonlogtest
}

func same(obj1, obj2 jsonlogtest) (string, bool) {
	diff := cmp.Diff(
		obj1,
		obj2,
		cmpopts.IgnoreFields(
			jsonlogtest{},
			"Timestamp",
		),
	)

	return diff, diff == ""
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
		panic("error marshaling jsonlogtest")
	} else {
		return string(out)
	}
}

func newjsonlog(
	pfix string,
	ltype string,
	err string,
	msgs ...string,
) jsonlogtest {
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
