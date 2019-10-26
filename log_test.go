package alog

import (
	"reflect"
	"testing"
	"time"
)

func Test_log_MarshalJSON(t *testing.T) {
	type fields struct {
		logger     *alog
		logtype    int8
		customtype string
		timestamp  time.Time
		err        error
		values     []interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := log{
				logger:     tt.fields.logger,
				logtype:    tt.fields.logtype,
				customtype: tt.fields.customtype,
				timestamp:  tt.fields.timestamp,
				err:        tt.fields.err,
				values:     tt.fields.values,
			}
			got, err := l.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("log.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("log.MarshalJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_log_Type(t *testing.T) {
	type fields struct {
		logger     *alog
		logtype    int8
		customtype string
		timestamp  time.Time
		err        error
		values     []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		wantT  string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := log{
				logger:     tt.fields.logger,
				logtype:    tt.fields.logtype,
				customtype: tt.fields.customtype,
				timestamp:  tt.fields.timestamp,
				err:        tt.fields.err,
				values:     tt.fields.values,
			}
			if gotT := l.Type(); gotT != tt.wantT {
				t.Errorf("log.Type() = %v, want %v", gotT, tt.wantT)
			}
		})
	}
}

func Test_log_String(t *testing.T) {
	type fields struct {
		logger     *alog
		logtype    int8
		customtype string
		timestamp  time.Time
		err        error
		values     interface{}
	}
	tests := []struct {
		name       string
		fields     fields
		wantOutput string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := log{
				logger:     tt.fields.logger,
				logtype:    tt.fields.logtype,
				customtype: tt.fields.customtype,
				timestamp:  tt.fields.timestamp,
				err:        tt.fields.err,
				values:     tt.fields.values,
			}
			if gotOutput := l.String(); gotOutput != tt.wantOutput {
				t.Errorf("log.String() = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}

func Test_log_getmessages(t *testing.T) {
	type fields struct {
		logger     *alog
		logtype    int8
		customtype string
		timestamp  time.Time
		err        error
		values     interface{}
	}
	type args struct {
		v interface{}
	}
	tests := []struct {
		name         string
		fields       fields
		args         args
		wantMessages []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := log{
				logger:     tt.fields.logger,
				logtype:    tt.fields.logtype,
				customtype: tt.fields.customtype,
				timestamp:  tt.fields.timestamp,
				err:        tt.fields.err,
				values:     tt.fields.values,
			}
			if gotMessages := l.getmessages(tt.args.v); !reflect.DeepEqual(gotMessages, tt.wantMessages) {
				t.Errorf("log.getmessages() = %v, want %v", gotMessages, tt.wantMessages)
			}
		})
	}
}

func Test_log_intslice(t *testing.T) {
	type fields struct {
		logger     *alog
		logtype    int8
		customtype string
		timestamp  time.Time
		err        error
		values     interface{}
	}
	type args struct {
		v []interface{}
	}
	tests := []struct {
		name          string
		fields        fields
		args          args
		wantFlattened []interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := log{
				logger:     tt.fields.logger,
				logtype:    tt.fields.logtype,
				customtype: tt.fields.customtype,
				timestamp:  tt.fields.timestamp,
				err:        tt.fields.err,
				values:     tt.fields.values,
			}
			if gotFlattened := l.intslice(tt.args.v); !reflect.DeepEqual(gotFlattened, tt.wantFlattened) {
				t.Errorf("log.intslice() = %v, want %v", gotFlattened, tt.wantFlattened)
			}
		})
	}
}

func Test_log_getmessage(t *testing.T) {
	type fields struct {
		logger     *alog
		logtype    int8
		customtype string
		timestamp  time.Time
		err        error
		values     interface{}
	}
	type args struct {
		v interface{}
	}
	tests := []struct {
		name        string
		fields      fields
		args        args
		wantMessage string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := log{
				logger:     tt.fields.logger,
				logtype:    tt.fields.logtype,
				customtype: tt.fields.customtype,
				timestamp:  tt.fields.timestamp,
				err:        tt.fields.err,
				values:     tt.fields.values,
			}
			if gotMessage := l.getmessage(tt.args.v); gotMessage != tt.wantMessage {
				t.Errorf("log.getmessage() = %v, want %v", gotMessage, tt.wantMessage)
			}
		})
	}
}
