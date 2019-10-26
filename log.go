package alog

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"time"
)

type log struct {
	logger     *alog
	logtype    int8
	customtype string
	timestamp  time.Time
	err        error
	values     interface{}
}

func (l log) String() (output string) {

	err := ""
	if l.err != nil {
		err = fmt.Sprintf(" %s", l.err)
	}

	strs := l.getmessages(l.values)

	output = fmt.Sprintf("%s %s%s %s\n",
		l.timestamp.Format(l.logger.dateformat),
		l.Type(),
		err,
		strings.Join(strs, ","),
	)

	return output
}

// getmessages breaks down the interface values and makes them
// into string messages that can then be represented in the different
// logging systems
func (l log) getmessages(v interface{}) (messages []string) {
	messages = make([]string, 0)

	switch x := l.values.(type) {
	case string:
		messages = append(messages, x)
	case []string:
		messages = append(messages, x...)
	case []interface{}:
		for _, val := range l.intslice(x) {
			messages = append(messages, l.getmessage(val))
		}
	default:
		messages = append(messages, l.getmessage(x))
	}

	return messages
}

// intslice takes an interface slice which may contain additional interface
// slices it to a singular interface slice so that it can be properly formatted
// by the logger
func (l log) intslice(v []interface{}) (flattened []interface{}) {

	flattened = make([]interface{}, 0)
	for _, value := range v {
		switch nv := value.(type) {
		case []interface{}:
			flattened = append(flattened, l.intslice(nv)...)
		default:
			flattened = append(flattened, nv)
		}
	}

	return flattened
}

// getmessage type switches the interface coming in to get a proper
// string value from each type based on the type selection
func (l log) getmessage(v interface{}) (message string) {

	switch field := v.(type) {
	case string:
		message = string(field)
	case fmt.Stringer:
		message = field.String()
	default:
		fmt.Println(reflect.TypeOf(field))
		message = fmt.Sprintf("%v", field)
	}

	return message
}

// MarshalJSON is used by the json marshaller to properly break
// down a log into a json struct for simpler parsing
func (l log) MarshalJSON() ([]byte, error) {

	// Setup a new flattened struct for json dumps of the logs
	output := &struct {
		LogType   string   `json:"type"`
		Timestamp string   `json:"timestamp"`
		Error     error    `json:"error,omitempty"`
		Messages  []string `json:"messages"`
	}{
		l.Type(),
		l.timestamp.Format(l.logger.dateformat),
		l.err,
		l.getmessages(l.values),
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
