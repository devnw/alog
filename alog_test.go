package alog

import (
	"bytes"
	"context"
	"io"
	"testing"
	"time"
)

func Test_alog_init(t *testing.T) {
	type fields struct {
		ctx        context.Context
		cancel     context.CancelFunc
		outputs    []io.Writer
		location   *time.Location
		dateformat string
		format     int
		prefix     string
		debug      bool
		feed       chan log
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &alog{
				ctx:        tt.fields.ctx,
				cancel:     tt.fields.cancel,
				outputs:    tt.fields.outputs,
				location:   tt.fields.location,
				dateformat: tt.fields.dateformat,
				format:     tt.fields.format,
				prefix:     tt.fields.prefix,
				debug:      tt.fields.debug,
				feed:       tt.fields.feed,
			}
			if err := l.init(); (err != nil) != tt.wantErr {
				t.Errorf("alog.init() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_alog_Print(t *testing.T) {
	type fields struct {
		ctx        context.Context
		cancel     context.CancelFunc
		outputs    []io.Writer
		location   *time.Location
		dateformat string
		format     int
		prefix     string
		debug      bool
		feed       chan log
	}
	type args struct {
		v []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &alog{
				ctx:        tt.fields.ctx,
				cancel:     tt.fields.cancel,
				outputs:    tt.fields.outputs,
				location:   tt.fields.location,
				dateformat: tt.fields.dateformat,
				format:     tt.fields.format,
				prefix:     tt.fields.prefix,
				debug:      tt.fields.debug,
				feed:       tt.fields.feed,
			}
			l.Print(tt.args.v...)
		})
	}
}

func Test_alog_Println(t *testing.T) {
	type fields struct {
		ctx        context.Context
		cancel     context.CancelFunc
		outputs    []io.Writer
		location   *time.Location
		dateformat string
		format     int
		prefix     string
		debug      bool
		feed       chan log
	}
	type args struct {
		v []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &alog{
				ctx:        tt.fields.ctx,
				cancel:     tt.fields.cancel,
				outputs:    tt.fields.outputs,
				location:   tt.fields.location,
				dateformat: tt.fields.dateformat,
				format:     tt.fields.format,
				prefix:     tt.fields.prefix,
				debug:      tt.fields.debug,
				feed:       tt.fields.feed,
			}
			l.Println(tt.args.v...)
		})
	}
}

func Test_alog_Printf(t *testing.T) {
	type fields struct {
		ctx        context.Context
		cancel     context.CancelFunc
		outputs    []io.Writer
		location   *time.Location
		dateformat string
		format     int
		prefix     string
		debug      bool
		feed       chan log
	}
	type args struct {
		format string
		v      []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &alog{
				ctx:        tt.fields.ctx,
				cancel:     tt.fields.cancel,
				outputs:    tt.fields.outputs,
				location:   tt.fields.location,
				dateformat: tt.fields.dateformat,
				format:     tt.fields.format,
				prefix:     tt.fields.prefix,
				debug:      tt.fields.debug,
				feed:       tt.fields.feed,
			}
			l.Printf(tt.args.format, tt.args.v...)
		})
	}
}

func Test_alog_Debug(t *testing.T) {
	type fields struct {
		ctx        context.Context
		cancel     context.CancelFunc
		outputs    []io.Writer
		location   *time.Location
		dateformat string
		format     int
		prefix     string
		debug      bool
		feed       chan log
	}
	type args struct {
		v []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &alog{
				ctx:        tt.fields.ctx,
				cancel:     tt.fields.cancel,
				outputs:    tt.fields.outputs,
				location:   tt.fields.location,
				dateformat: tt.fields.dateformat,
				format:     tt.fields.format,
				prefix:     tt.fields.prefix,
				debug:      tt.fields.debug,
				feed:       tt.fields.feed,
			}
			l.Debug(tt.args.v...)
		})
	}
}

func Test_alog_Debugln(t *testing.T) {
	type fields struct {
		ctx        context.Context
		cancel     context.CancelFunc
		outputs    []io.Writer
		location   *time.Location
		dateformat string
		format     int
		prefix     string
		debug      bool
		feed       chan log
	}
	type args struct {
		v []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &alog{
				ctx:        tt.fields.ctx,
				cancel:     tt.fields.cancel,
				outputs:    tt.fields.outputs,
				location:   tt.fields.location,
				dateformat: tt.fields.dateformat,
				format:     tt.fields.format,
				prefix:     tt.fields.prefix,
				debug:      tt.fields.debug,
				feed:       tt.fields.feed,
			}
			l.Debugln(tt.args.v...)
		})
	}
}

func Test_alog_Debugf(t *testing.T) {
	type fields struct {
		ctx        context.Context
		cancel     context.CancelFunc
		outputs    []io.Writer
		location   *time.Location
		dateformat string
		format     int
		prefix     string
		debug      bool
		feed       chan log
	}
	type args struct {
		format string
		v      []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &alog{
				ctx:        tt.fields.ctx,
				cancel:     tt.fields.cancel,
				outputs:    tt.fields.outputs,
				location:   tt.fields.location,
				dateformat: tt.fields.dateformat,
				format:     tt.fields.format,
				prefix:     tt.fields.prefix,
				debug:      tt.fields.debug,
				feed:       tt.fields.feed,
			}
			l.Debugf(tt.args.format, tt.args.v...)
		})
	}
}

func Test_alog_Warn(t *testing.T) {
	type fields struct {
		ctx        context.Context
		cancel     context.CancelFunc
		outputs    []io.Writer
		location   *time.Location
		dateformat string
		format     int
		prefix     string
		debug      bool
		feed       chan log
	}
	type args struct {
		err error
		v   []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &alog{
				ctx:        tt.fields.ctx,
				cancel:     tt.fields.cancel,
				outputs:    tt.fields.outputs,
				location:   tt.fields.location,
				dateformat: tt.fields.dateformat,
				format:     tt.fields.format,
				prefix:     tt.fields.prefix,
				debug:      tt.fields.debug,
				feed:       tt.fields.feed,
			}
			l.Warn(tt.args.err, tt.args.v...)
		})
	}
}

func Test_alog_Warnln(t *testing.T) {
	type fields struct {
		ctx        context.Context
		cancel     context.CancelFunc
		outputs    []io.Writer
		location   *time.Location
		dateformat string
		format     int
		prefix     string
		debug      bool
		feed       chan log
	}
	type args struct {
		err error
		v   []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &alog{
				ctx:        tt.fields.ctx,
				cancel:     tt.fields.cancel,
				outputs:    tt.fields.outputs,
				location:   tt.fields.location,
				dateformat: tt.fields.dateformat,
				format:     tt.fields.format,
				prefix:     tt.fields.prefix,
				debug:      tt.fields.debug,
				feed:       tt.fields.feed,
			}
			l.Warnln(tt.args.err, tt.args.v...)
		})
	}
}

func Test_alog_Warnf(t *testing.T) {
	type fields struct {
		ctx        context.Context
		cancel     context.CancelFunc
		outputs    []io.Writer
		location   *time.Location
		dateformat string
		format     int
		prefix     string
		debug      bool
		feed       chan log
	}
	type args struct {
		err    error
		format string
		v      []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &alog{
				ctx:        tt.fields.ctx,
				cancel:     tt.fields.cancel,
				outputs:    tt.fields.outputs,
				location:   tt.fields.location,
				dateformat: tt.fields.dateformat,
				format:     tt.fields.format,
				prefix:     tt.fields.prefix,
				debug:      tt.fields.debug,
				feed:       tt.fields.feed,
			}
			l.Warnf(tt.args.err, tt.args.format, tt.args.v...)
		})
	}
}

func Test_alog_Error(t *testing.T) {
	type fields struct {
		ctx        context.Context
		cancel     context.CancelFunc
		outputs    []io.Writer
		location   *time.Location
		dateformat string
		format     int
		prefix     string
		debug      bool
		feed       chan log
	}
	type args struct {
		err error
		v   []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &alog{
				ctx:        tt.fields.ctx,
				cancel:     tt.fields.cancel,
				outputs:    tt.fields.outputs,
				location:   tt.fields.location,
				dateformat: tt.fields.dateformat,
				format:     tt.fields.format,
				prefix:     tt.fields.prefix,
				debug:      tt.fields.debug,
				feed:       tt.fields.feed,
			}
			l.Error(tt.args.err, tt.args.v...)
		})
	}
}

func Test_alog_Errorln(t *testing.T) {
	type fields struct {
		ctx        context.Context
		cancel     context.CancelFunc
		outputs    []io.Writer
		location   *time.Location
		dateformat string
		format     int
		prefix     string
		debug      bool
		feed       chan log
	}
	type args struct {
		err error
		v   []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &alog{
				ctx:        tt.fields.ctx,
				cancel:     tt.fields.cancel,
				outputs:    tt.fields.outputs,
				location:   tt.fields.location,
				dateformat: tt.fields.dateformat,
				format:     tt.fields.format,
				prefix:     tt.fields.prefix,
				debug:      tt.fields.debug,
				feed:       tt.fields.feed,
			}
			l.Errorln(tt.args.err, tt.args.v...)
		})
	}
}

func Test_alog_Errorf(t *testing.T) {
	type fields struct {
		ctx        context.Context
		cancel     context.CancelFunc
		outputs    []io.Writer
		location   *time.Location
		dateformat string
		format     int
		prefix     string
		debug      bool
		feed       chan log
	}
	type args struct {
		err    error
		format string
		v      []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &alog{
				ctx:        tt.fields.ctx,
				cancel:     tt.fields.cancel,
				outputs:    tt.fields.outputs,
				location:   tt.fields.location,
				dateformat: tt.fields.dateformat,
				format:     tt.fields.format,
				prefix:     tt.fields.prefix,
				debug:      tt.fields.debug,
				feed:       tt.fields.feed,
			}
			l.Errorf(tt.args.err, tt.args.format, tt.args.v...)
		})
	}
}

func Test_alog_Crit(t *testing.T) {
	type fields struct {
		ctx        context.Context
		cancel     context.CancelFunc
		outputs    []io.Writer
		location   *time.Location
		dateformat string
		format     int
		prefix     string
		debug      bool
		feed       chan log
	}
	type args struct {
		err error
		v   []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &alog{
				ctx:        tt.fields.ctx,
				cancel:     tt.fields.cancel,
				outputs:    tt.fields.outputs,
				location:   tt.fields.location,
				dateformat: tt.fields.dateformat,
				format:     tt.fields.format,
				prefix:     tt.fields.prefix,
				debug:      tt.fields.debug,
				feed:       tt.fields.feed,
			}
			l.Crit(tt.args.err, tt.args.v...)
		})
	}
}

func Test_alog_Critln(t *testing.T) {
	type fields struct {
		ctx        context.Context
		cancel     context.CancelFunc
		outputs    []io.Writer
		location   *time.Location
		dateformat string
		format     int
		prefix     string
		debug      bool
		feed       chan log
	}
	type args struct {
		err error
		v   []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &alog{
				ctx:        tt.fields.ctx,
				cancel:     tt.fields.cancel,
				outputs:    tt.fields.outputs,
				location:   tt.fields.location,
				dateformat: tt.fields.dateformat,
				format:     tt.fields.format,
				prefix:     tt.fields.prefix,
				debug:      tt.fields.debug,
				feed:       tt.fields.feed,
			}
			l.Critln(tt.args.err, tt.args.v...)
		})
	}
}

func Test_alog_Critf(t *testing.T) {
	type fields struct {
		ctx        context.Context
		cancel     context.CancelFunc
		outputs    []io.Writer
		location   *time.Location
		dateformat string
		format     int
		prefix     string
		debug      bool
		feed       chan log
	}
	type args struct {
		err    error
		format string
		v      []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &alog{
				ctx:        tt.fields.ctx,
				cancel:     tt.fields.cancel,
				outputs:    tt.fields.outputs,
				location:   tt.fields.location,
				dateformat: tt.fields.dateformat,
				format:     tt.fields.format,
				prefix:     tt.fields.prefix,
				debug:      tt.fields.debug,
				feed:       tt.fields.feed,
			}
			l.Critf(tt.args.err, tt.args.format, tt.args.v...)
		})
	}
}

func Test_alog_Fatal(t *testing.T) {
	type fields struct {
		ctx        context.Context
		cancel     context.CancelFunc
		outputs    []io.Writer
		location   *time.Location
		dateformat string
		format     int
		prefix     string
		debug      bool
		feed       chan log
	}
	type args struct {
		err error
		v   []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &alog{
				ctx:        tt.fields.ctx,
				cancel:     tt.fields.cancel,
				outputs:    tt.fields.outputs,
				location:   tt.fields.location,
				dateformat: tt.fields.dateformat,
				format:     tt.fields.format,
				prefix:     tt.fields.prefix,
				debug:      tt.fields.debug,
				feed:       tt.fields.feed,
			}
			l.Fatal(tt.args.err, tt.args.v...)
		})
	}
}

func Test_alog_Fatalln(t *testing.T) {
	type fields struct {
		ctx        context.Context
		cancel     context.CancelFunc
		outputs    []io.Writer
		location   *time.Location
		dateformat string
		format     int
		prefix     string
		debug      bool
		feed       chan log
	}
	type args struct {
		err error
		v   []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &alog{
				ctx:        tt.fields.ctx,
				cancel:     tt.fields.cancel,
				outputs:    tt.fields.outputs,
				location:   tt.fields.location,
				dateformat: tt.fields.dateformat,
				format:     tt.fields.format,
				prefix:     tt.fields.prefix,
				debug:      tt.fields.debug,
				feed:       tt.fields.feed,
			}
			l.Fatalln(tt.args.err, tt.args.v...)
		})
	}
}

func Test_alog_Fatalf(t *testing.T) {
	type fields struct {
		ctx        context.Context
		cancel     context.CancelFunc
		outputs    []io.Writer
		location   *time.Location
		dateformat string
		format     int
		prefix     string
		debug      bool
		feed       chan log
	}
	type args struct {
		err    error
		format string
		v      []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &alog{
				ctx:        tt.fields.ctx,
				cancel:     tt.fields.cancel,
				outputs:    tt.fields.outputs,
				location:   tt.fields.location,
				dateformat: tt.fields.dateformat,
				format:     tt.fields.format,
				prefix:     tt.fields.prefix,
				debug:      tt.fields.debug,
				feed:       tt.fields.feed,
			}
			l.Fatalf(tt.args.err, tt.args.format, tt.args.v...)
		})
	}
}

func Test_alog_AddOutput(t *testing.T) {
	type fields struct {
		ctx        context.Context
		cancel     context.CancelFunc
		outputs    []io.Writer
		location   *time.Location
		dateformat string
		format     int
		prefix     string
		debug      bool
		feed       chan log
	}
	tests := []struct {
		name    string
		fields  fields
		wantOut string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &alog{
				ctx:        tt.fields.ctx,
				cancel:     tt.fields.cancel,
				outputs:    tt.fields.outputs,
				location:   tt.fields.location,
				dateformat: tt.fields.dateformat,
				format:     tt.fields.format,
				prefix:     tt.fields.prefix,
				debug:      tt.fields.debug,
				feed:       tt.fields.feed,
			}
			out := &bytes.Buffer{}
			l.AddOutput(out)
			if gotOut := out.String(); gotOut != tt.wantOut {
				t.Errorf("alog.AddOutput() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}

func Test_alog_Close(t *testing.T) {
	type fields struct {
		ctx        context.Context
		cancel     context.CancelFunc
		outputs    []io.Writer
		location   *time.Location
		dateformat string
		format     int
		prefix     string
		debug      bool
		feed       chan log
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &alog{
				ctx:        tt.fields.ctx,
				cancel:     tt.fields.cancel,
				outputs:    tt.fields.outputs,
				location:   tt.fields.location,
				dateformat: tt.fields.dateformat,
				format:     tt.fields.format,
				prefix:     tt.fields.prefix,
				debug:      tt.fields.debug,
				feed:       tt.fields.feed,
			}
			l.Close()
		})
	}
}

func Test_alog_Validate(t *testing.T) {
	type fields struct {
		ctx        context.Context
		cancel     context.CancelFunc
		outputs    []io.Writer
		location   *time.Location
		dateformat string
		format     int
		prefix     string
		debug      bool
		feed       chan log
	}
	tests := []struct {
		name      string
		fields    fields
		wantValid bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &alog{
				ctx:        tt.fields.ctx,
				cancel:     tt.fields.cancel,
				outputs:    tt.fields.outputs,
				location:   tt.fields.location,
				dateformat: tt.fields.dateformat,
				format:     tt.fields.format,
				prefix:     tt.fields.prefix,
				debug:      tt.fields.debug,
				feed:       tt.fields.feed,
			}
			if gotValid := l.Validate(); gotValid != tt.wantValid {
				t.Errorf("alog.Validate() = %v, want %v", gotValid, tt.wantValid)
			}
		})
	}
}
