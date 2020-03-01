package alog

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type log struct {
	logger     *alog
	logtype    LogLevel
	customtype string
	timestamp  time.Time
	err        error
	values     []interface{}
}

func (l log) String() (output string) {

	err := ""
	if l.err != nil {
		err = fmt.Sprintf(" | err: %s", l.err.Error())
	}

	message := ""
	strs := l.getmessages(l.values)
	if len(strs) > 0 {
		message = fmt.Sprintf(" %s", strings.Join(strs, ","))
	}

	prefix := ""
	if len(l.logger.prefix) > 0 {
		prefix = fmt.Sprintf("%s ", l.logger.prefix)
	}

	// Handle the empty message and error section
	if len(err) == 0 && len(message) == 0 {
		message = "unable to create log string, empty message and error"
	}

	output = fmt.Sprintf("%s%s [%s]%s%s",
		prefix,
		l.timestamp.Format(l.logger.dateformat),
		l.Type(),
		message,
		err,
	)

	if string(output[len(output)-1]) != "\n" {
		output = fmt.Sprintf("%s\n", output)
	}

	return output
}

// getmessages breaks down the interface values and makes them
// into string messages that can then be represented in the different
// logging systems
func (l log) getmessages(v []interface{}) (messages []string) {
	messages = make([]string, 0)

	for _, val := range l.intslice(v) {
		messages = append(messages, l.getmessage(val))
	}

	return messages
}

// intslice takes an interface slice which may contain additional interface
// slices it to a singular interface slice so that it can be properly formatted
// by the logger
func (l log) intslice(v []interface{}) (flattened []interface{}) {

	flattened = make([]interface{}, 0)
	for _, value := range v {

		switch x := value.(type) {
		case string:
			flattened = append(flattened, x)
		case []string:
			for _, s := range x {
				flattened = append(flattened, s)
			}
		case []interface{}:
			flattened = append(flattened, l.intslice(x)...)
		default:
			flattened = append(flattened, l.getmessage(x))
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
		message = fmt.Sprintf("%v", field)
	}

	return message
}

// MarshalJSON is used by the json marshaller to properly break
// down a log into a json struct for simpler parsing
func (l log) MarshalJSON() ([]byte, error) {

	var err *string
	if l.err != nil {
		e := l.err.Error()
		err = &e
	}

	var prefix *string
	if len(l.logger.prefix) > 0 {
		prefix = &l.logger.prefix
	}

	// Setup a new flattened struct for json dumps of the logs
	output := &struct {
		Prefix    *string  `json:"prefix,omitempty"`
		LogType   string   `json:"type"`
		Timestamp string   `json:"timestamp"`
		Error     *string  `json:"error,omitempty"`
		Messages  []string `json:"messages"`
	}{
		prefix,
		l.Type(),
		l.timestamp.Format(l.logger.dateformat),
		err,
		l.getmessages(l.values),
	}

	return json.Marshal(output)
}

// Type returns the type of the log for parsing or displaying
func (l log) Type() (t string) {

	if l.logtype&CUSTOM > 0 {
		t = "CUSTOM"
		if len(l.customtype) > 0 {
			t = strings.ToUpper(l.customtype)
		}
	} else {
		var ok bool
		if t, ok = types[l.logtype]; !ok {

			// Use INFO as default type
			t = types[INFO]
		}
	}

	return t
}
