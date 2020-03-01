package alog

import "github.com/pkg/errors"

var logsDELIM = []fakelog{
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

var cerrlogsDELIM = []fakelog{
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

var clogsDELIM = []fakelog{
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

var multiDELIM = []fakelog{
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

var flogsDELIM = []fakelog{
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
