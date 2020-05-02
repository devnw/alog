package alog

import (
	"context"
	"io"
	"os"
	"testing"

	"github.com/pkg/errors"
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

// TestDestinations returns a list of destinations for logging test data
// to the *testing.T that was passed in. This defaults all ERROR, CRIT, FATAL
// logs to the t.Error and the rest are routed to t.Log. These destinations
// can be used to override the logger with destinations specific to testing.
func TestDestinations(ctx context.Context, t *testing.T) []Destination {

	return []Destination{
		{
			INFO | DEBUG | TRACE | WARN | CUSTOM,
			STD,
			test{
				ctx:  ctx,
				t:    t,
				mode: INFO | DEBUG | TRACE | WARN | CUSTOM,
			},
		},
		{
			ERROR | CRIT | FATAL,
			STD,
			test{
				ctx:  ctx,
				t:    t,
				mode: ERROR | CRIT | FATAL,
			},
		},
	}
}

type test struct {
	ctx  context.Context
	t    *testing.T
	mode LogLevel
}

func (t test) Write(p []byte) (int, error) {
	if t.t == nil {
		return 0, errors.New("invalid test object")
	}

	select {
	case <-t.ctx.Done():
		return 0, nil
	default:
		msg := string(p)
		if t.mode&(ERROR|CRIT|FATAL) > 0 {
			t.t.Error(msg)
		} else {
			t.t.Log(msg)
		}
	}

	return len(p), nil
}

// BenchDestinations returns a list of destinations for benchmarking.
// The destination returned from this method does NOTHING. It is meant
// to remove overhead from the logger for proper benchmarks.
// This destination can be used to override the logger for benchmarking.
func BenchDestinations() []Destination {
	return []Destination{
		{
			INFO | DEBUG | TRACE | WARN |
				CUSTOM | ERROR | CRIT | FATAL,
			STD,
			bench{},
		},
	}
}

type bench struct {
}

func (bench) Write(p []byte) (int, error) {
	return len(p), nil
}
