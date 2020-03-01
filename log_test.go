package alog

import "testing"

func Test_log_Type(t *testing.T) {
	l := log{
		logtype: LogLevel(1555),
	}

	if l.Type() != "INFO" {
		t.Error("unknown type is mis-matched in logger")
	}
}
