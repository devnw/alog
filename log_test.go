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
