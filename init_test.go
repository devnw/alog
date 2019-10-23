package alog

import (
	"context"
	"io"
	"reflect"
	"testing"
	"time"
)

func Test_setGlobal(t *testing.T) {
	type args struct {
		logger Logger
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := setGlobal(tt.args.logger); (err != nil) != tt.wantErr {
				t.Errorf("setGlobal() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewGlobal(t *testing.T) {
	type args struct {
		ctx        context.Context
		format     int
		prefix     string
		dateformat string
		location   *time.Location
		debug      bool
		buffer     int
		out        []io.Writer
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := NewGlobal(tt.args.ctx, tt.args.format, tt.args.prefix, tt.args.dateformat, tt.args.location, tt.args.debug, tt.args.buffer, tt.args.out...); (err != nil) != tt.wantErr {
				t.Errorf("NewGlobal() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNew(t *testing.T) {
	type args struct {
		ctx        context.Context
		format     int
		prefix     string
		dateformat string
		location   *time.Location
		debug      bool
		buffer     int
		out        []io.Writer
	}
	tests := []struct {
		name       string
		args       args
		wantLogger Logger
		wantErr    bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotLogger, err := New(tt.args.ctx, tt.args.format, tt.args.prefix, tt.args.dateformat, tt.args.location, tt.args.debug, tt.args.buffer, tt.args.out...)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotLogger, tt.wantLogger) {
				t.Errorf("New() = %v, want %v", gotLogger, tt.wantLogger)
			}
		})
	}
}
