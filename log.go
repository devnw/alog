package alog

import (
	"encoding/json"
	"fmt"
	"time"
)

type log struct {
	logger     *alog
	logtype    int8
	customtype string
	timestamp  time.Time
	err        error
	format     *string
	v          []interface{}
}

func (l log) MarshalJSON() ([]byte, error) {

	// Setup the messages for the marshalling
	var messages []interface{}
	if l.format != nil {
		messages = append(messages, fmt.Sprintf(*l.format, l.v))
	} else {
		messages = append(messages, l.v...)
	}

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
		messages,
	}

	return json.Marshal(output)
}

// Type returns the type of the log for parsing or displaying
func (l log) Type() (t string) {

	if l.logtype&FATAL == 1 {
		t = "FATAL"
	} else if l.logtype&CRIT == 1 {
		t = "CRITICAL"
	} else if l.logtype&ERROR == 1 {
		t = "ERROR"
	} else if l.logtype&WARN == 1 {
		t = "WARNING"
	} else if l.logtype&DEBUG == 1 {
		t = "DEBUG"
	} else if l.logtype&CUSTOM == 1 {
		t = "CUSTOM"
		if len(l.customtype) > 0 {
			t = l.customtype
		}
	} else {
		t = "INFO"
	}

	return t
}
