package alog

import (
	"io"
	"os"
)

// Dest is the destination struct for registering io.Writers to the
// alog library so that different log types can be passed to each writer
// asynchronously
type Dest struct {
	Types  int8
	Format int8
	Writer io.Writer
}

// Standards returns the standard out and standard error destinations for
// quick access when creating a logger
// INFO, DEBUG, WARNING Logs are logged to Standard Out
// ERROR, CRITICAL, FATAL Logs are logged to Standard Error
func Standards() []Dest {
	return []Dest{
		{
			INFO | DEBUG | WARN,
			STD,
			os.Stdout,
		},
		{
			ERROR | CRIT | FATAL,
			STD,
			os.Stderr,
		},
	}
}
