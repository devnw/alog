package alog

import (
	"io"
	"os"
)

// Destination is the destination struct for registering io.Writers to the
// alog library so that different log types can be passed to each writer
// asynchronously
type Destination struct {
	Types  LogLevel
	Format LogFmt
	Writer io.Writer
}

// Standards returns the standard out and standard error destinations for
// quick access when creating a logger
// INFO, DEBUG, TRACE, WARNING, CUSTOM Logs are logged to Standard Out
// ERROR, CRITICAL, FATAL Logs are logged to Standard Error
func Standards() []Destination {

	return []Destination{
		{
			INFO | DEBUG | TRACE | WARN | CUSTOM,
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
