package alog

import (
	"encoding/json"
	"time"
)

type log struct {
	logger     *alog
	logtype    int8
	customtype string
	timestamp  time.Time
	err        error
	values     []interface{}
}

func (l log) MarshalJSON() ([]byte, error) {

	// Setup a new flattened struct for json dumps of the logs
	output := &struct {
		LogType   string        `json:"type"`
		Timestamp string        `json:"timestamp"`
		Error     error         `json:"error,omitempty"`
		Messages  []interface{} `json:"messages"`
	}{
		l.Type(),
		l.timestamp.Format(l.logger.dateformat),
		l.err,
		l.values,
	}

	return json.Marshal(output)
}

// Type returns the type of the log for parsing or displaying
func (l log) Type() (t string) {

	if l.logtype&FATAL > 0 {
		t = "FATAL"
	} else if l.logtype&CRIT > 0 {
		t = "CRITICAL"
	} else if l.logtype&ERROR > 0 {
		t = "ERROR"
	} else if l.logtype&WARN > 0 {
		t = "WARNING"
	} else if l.logtype&DEBUG > 0 {
		t = "DEBUG"
	} else if l.logtype&CUSTOM > 0 {
		t = "CUSTOM"
		if len(l.customtype) > 0 {
			t = l.customtype
		}
	} else {
		t = "INFO"
	}

	return t
}
