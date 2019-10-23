package alog

type log struct {
	err    error
	format *string
	v      []interface{}
}
