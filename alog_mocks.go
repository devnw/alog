package alog

import (
	"github.com/pkg/errors"
)

type passmock struct {
	msg chan []byte
}

func (pm *passmock) Write(p []byte) (n int, err error) {

	n = len(p)
	pm.msg <- p

	return n, err
}

type fakelog struct {
	lvl      LogLevel
	text     string
	err      error
	expected string
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
