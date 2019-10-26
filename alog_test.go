package alog

import (
	"bytes"
	"context"
	"reflect"
	"testing"
	"time"
)

func Test_alog_global(t *testing.T) {
	Println("HELLO WORLD")
}

func Test_alog_init(t *testing.T) {
	type fields struct {
		ctx          context.Context
		cancel       context.CancelFunc
		destinations []Destination
		location     *time.Location
		dateformat   string
		prefix       string
		logdebug     bool
		buffer       int
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
				ctx:          tt.fields.ctx,
				cancel:       tt.fields.cancel,
				destinations: tt.fields.destinations,
				location:     tt.fields.location,
				dateformat:   tt.fields.dateformat,
				prefix:       tt.fields.prefix,
				logdebug:     tt.fields.logdebug,
				buffer:       tt.fields.buffer,
			}
			if err := l.init(); (err != nil) != tt.wantErr {
				t.Errorf("alog.init() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_alog_listen(t *testing.T) {
	type fields struct {
		ctx          context.Context
		cancel       context.CancelFunc
		destinations []Destination
		location     *time.Location
		dateformat   string
		prefix       string
		logdebug     bool
		buffer       int
	}
	type args struct {
		ctx         context.Context
		destination Destination
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   chan<- log
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &alog{
				ctx:          tt.fields.ctx,
				cancel:       tt.fields.cancel,
				destinations: tt.fields.destinations,
				location:     tt.fields.location,
				dateformat:   tt.fields.dateformat,
				prefix:       tt.fields.prefix,
				logdebug:     tt.fields.logdebug,
				buffer:       tt.fields.buffer,
			}
			if got := l.listen(tt.args.ctx, tt.args.destination); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("alog.listen() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_alog_send(t *testing.T) {
	type fields struct {
		ctx          context.Context
		cancel       context.CancelFunc
		destinations []Destination
		location     *time.Location
		dateformat   string
		prefix       string
		logdebug     bool
		buffer       int
	}
	type args struct {
		ctx     context.Context
		logtype int8
		value   log
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
				ctx:          tt.fields.ctx,
				cancel:       tt.fields.cancel,
				destinations: tt.fields.destinations,
				location:     tt.fields.location,
				dateformat:   tt.fields.dateformat,
				prefix:       tt.fields.prefix,
				logdebug:     tt.fields.logdebug,
				buffer:       tt.fields.buffer,
			}
			l.send(tt.args.ctx, tt.args.value)
		})
	}
}

func Test_alog_getd(t *testing.T) {
	type fields struct {
		ctx          context.Context
		cancel       context.CancelFunc
		destinations []Destination
		location     *time.Location
		dateformat   string
		prefix       string
		logdebug     bool
		buffer       int
	}
	type args struct {
		logtype int8
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []chan<- log
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &alog{
				ctx:          tt.fields.ctx,
				cancel:       tt.fields.cancel,
				destinations: tt.fields.destinations,
				location:     tt.fields.location,
				dateformat:   tt.fields.dateformat,
				prefix:       tt.fields.prefix,
				logdebug:     tt.fields.logdebug,
				buffer:       tt.fields.buffer,
			}
			if got := l.getd(tt.args.logtype); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("alog.getd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_alog_Print(t *testing.T) {
	type fields struct {
		ctx          context.Context
		cancel       context.CancelFunc
		destinations []Destination
		location     *time.Location
		dateformat   string
		prefix       string
		logdebug     bool
		buffer       int
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
				ctx:          tt.fields.ctx,
				cancel:       tt.fields.cancel,
				destinations: tt.fields.destinations,
				location:     tt.fields.location,
				dateformat:   tt.fields.dateformat,
				prefix:       tt.fields.prefix,
				logdebug:     tt.fields.logdebug,
				buffer:       tt.fields.buffer,
			}
			l.Print(tt.args.v...)
		})
	}
}

func Test_alog_Println(t *testing.T) {
	type fields struct {
		ctx          context.Context
		cancel       context.CancelFunc
		destinations []Destination
		location     *time.Location
		dateformat   string
		prefix       string
		logdebug     bool
		buffer       int
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
				ctx:          tt.fields.ctx,
				cancel:       tt.fields.cancel,
				destinations: tt.fields.destinations,
				location:     tt.fields.location,
				dateformat:   tt.fields.dateformat,
				prefix:       tt.fields.prefix,
				logdebug:     tt.fields.logdebug,
				buffer:       tt.fields.buffer,
			}
			l.Println(tt.args.v...)
		})
	}
}

func Test_alog_Printf(t *testing.T) {
	type fields struct {
		ctx          context.Context
		cancel       context.CancelFunc
		destinations []Destination
		location     *time.Location
		dateformat   string
		prefix       string
		logdebug     bool
		buffer       int
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
				ctx:          tt.fields.ctx,
				cancel:       tt.fields.cancel,
				destinations: tt.fields.destinations,
				location:     tt.fields.location,
				dateformat:   tt.fields.dateformat,
				prefix:       tt.fields.prefix,
				logdebug:     tt.fields.logdebug,
				buffer:       tt.fields.buffer,
			}
			l.Printf(tt.args.format, tt.args.v...)
		})
	}
}

func Test_alog_Debug(t *testing.T) {
	type fields struct {
		ctx          context.Context
		cancel       context.CancelFunc
		destinations []Destination
		location     *time.Location
		dateformat   string
		prefix       string
		logdebug     bool
		buffer       int
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
				ctx:          tt.fields.ctx,
				cancel:       tt.fields.cancel,
				destinations: tt.fields.destinations,
				location:     tt.fields.location,
				dateformat:   tt.fields.dateformat,
				prefix:       tt.fields.prefix,
				logdebug:     tt.fields.logdebug,
				buffer:       tt.fields.buffer,
			}
			l.Debug(tt.args.v...)
		})
	}
}

func Test_alog_Debugln(t *testing.T) {
	type fields struct {
		ctx          context.Context
		cancel       context.CancelFunc
		destinations []Destination
		location     *time.Location
		dateformat   string
		prefix       string
		logdebug     bool
		buffer       int
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
				ctx:          tt.fields.ctx,
				cancel:       tt.fields.cancel,
				destinations: tt.fields.destinations,
				location:     tt.fields.location,
				dateformat:   tt.fields.dateformat,
				prefix:       tt.fields.prefix,
				logdebug:     tt.fields.logdebug,
				buffer:       tt.fields.buffer,
			}
			l.Debugln(tt.args.v...)
		})
	}
}

func Test_alog_Debugf(t *testing.T) {
	type fields struct {
		ctx          context.Context
		cancel       context.CancelFunc
		destinations []Destination
		location     *time.Location
		dateformat   string
		prefix       string
		logdebug     bool
		buffer       int
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
				ctx:          tt.fields.ctx,
				cancel:       tt.fields.cancel,
				destinations: tt.fields.destinations,
				location:     tt.fields.location,
				dateformat:   tt.fields.dateformat,
				prefix:       tt.fields.prefix,
				logdebug:     tt.fields.logdebug,
				buffer:       tt.fields.buffer,
			}
			l.Debugf(tt.args.format, tt.args.v...)
		})
	}
}

func Test_alog_Warn(t *testing.T) {
	type fields struct {
		ctx          context.Context
		cancel       context.CancelFunc
		destinations []Destination
		location     *time.Location
		dateformat   string
		prefix       string
		logdebug     bool
		buffer       int
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
				ctx:          tt.fields.ctx,
				cancel:       tt.fields.cancel,
				destinations: tt.fields.destinations,
				location:     tt.fields.location,
				dateformat:   tt.fields.dateformat,
				prefix:       tt.fields.prefix,
				logdebug:     tt.fields.logdebug,
				buffer:       tt.fields.buffer,
			}
			l.Warn(tt.args.err, tt.args.v...)
		})
	}
}

func Test_alog_Warnln(t *testing.T) {
	type fields struct {
		ctx          context.Context
		cancel       context.CancelFunc
		destinations []Destination
		location     *time.Location
		dateformat   string
		prefix       string
		logdebug     bool
		buffer       int
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
				ctx:          tt.fields.ctx,
				cancel:       tt.fields.cancel,
				destinations: tt.fields.destinations,
				location:     tt.fields.location,
				dateformat:   tt.fields.dateformat,
				prefix:       tt.fields.prefix,
				logdebug:     tt.fields.logdebug,
				buffer:       tt.fields.buffer,
			}
			l.Warnln(tt.args.err, tt.args.v...)
		})
	}
}

func Test_alog_Warnf(t *testing.T) {
	type fields struct {
		ctx          context.Context
		cancel       context.CancelFunc
		destinations []Destination
		location     *time.Location
		dateformat   string
		prefix       string
		logdebug     bool
		buffer       int
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
				ctx:          tt.fields.ctx,
				cancel:       tt.fields.cancel,
				destinations: tt.fields.destinations,
				location:     tt.fields.location,
				dateformat:   tt.fields.dateformat,
				prefix:       tt.fields.prefix,
				logdebug:     tt.fields.logdebug,
				buffer:       tt.fields.buffer,
			}
			l.Warnf(tt.args.err, tt.args.format, tt.args.v...)
		})
	}
}

func Test_alog_Error(t *testing.T) {
	type fields struct {
		ctx          context.Context
		cancel       context.CancelFunc
		destinations []Destination
		location     *time.Location
		dateformat   string
		prefix       string
		logdebug     bool
		buffer       int
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
				ctx:          tt.fields.ctx,
				cancel:       tt.fields.cancel,
				destinations: tt.fields.destinations,
				location:     tt.fields.location,
				dateformat:   tt.fields.dateformat,
				prefix:       tt.fields.prefix,
				logdebug:     tt.fields.logdebug,
				buffer:       tt.fields.buffer,
			}
			l.Error(tt.args.err, tt.args.v...)
		})
	}
}

func Test_alog_Errorln(t *testing.T) {
	type fields struct {
		ctx          context.Context
		cancel       context.CancelFunc
		destinations []Destination
		location     *time.Location
		dateformat   string
		prefix       string
		logdebug     bool
		buffer       int
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
				ctx:          tt.fields.ctx,
				cancel:       tt.fields.cancel,
				destinations: tt.fields.destinations,
				location:     tt.fields.location,
				dateformat:   tt.fields.dateformat,
				prefix:       tt.fields.prefix,
				logdebug:     tt.fields.logdebug,
				buffer:       tt.fields.buffer,
			}
			l.Errorln(tt.args.err, tt.args.v...)
		})
	}
}

func Test_alog_Errorf(t *testing.T) {
	type fields struct {
		ctx          context.Context
		cancel       context.CancelFunc
		destinations []Destination
		location     *time.Location
		dateformat   string
		prefix       string
		logdebug     bool
		buffer       int
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
				ctx:          tt.fields.ctx,
				cancel:       tt.fields.cancel,
				destinations: tt.fields.destinations,
				location:     tt.fields.location,
				dateformat:   tt.fields.dateformat,
				prefix:       tt.fields.prefix,
				logdebug:     tt.fields.logdebug,
				buffer:       tt.fields.buffer,
			}
			l.Errorf(tt.args.err, tt.args.format, tt.args.v...)
		})
	}
}

func Test_alog_Crit(t *testing.T) {
	type fields struct {
		ctx          context.Context
		cancel       context.CancelFunc
		destinations []Destination
		location     *time.Location
		dateformat   string
		prefix       string
		logdebug     bool
		buffer       int
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
				ctx:          tt.fields.ctx,
				cancel:       tt.fields.cancel,
				destinations: tt.fields.destinations,
				location:     tt.fields.location,
				dateformat:   tt.fields.dateformat,
				prefix:       tt.fields.prefix,
				logdebug:     tt.fields.logdebug,
				buffer:       tt.fields.buffer,
			}
			l.Crit(tt.args.err, tt.args.v...)
		})
	}
}

func Test_alog_Critln(t *testing.T) {
	type fields struct {
		ctx          context.Context
		cancel       context.CancelFunc
		destinations []Destination
		location     *time.Location
		dateformat   string
		prefix       string
		logdebug     bool
		buffer       int
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
				ctx:          tt.fields.ctx,
				cancel:       tt.fields.cancel,
				destinations: tt.fields.destinations,
				location:     tt.fields.location,
				dateformat:   tt.fields.dateformat,
				prefix:       tt.fields.prefix,
				logdebug:     tt.fields.logdebug,
				buffer:       tt.fields.buffer,
			}
			l.Critln(tt.args.err, tt.args.v...)
		})
	}
}

func Test_alog_Critf(t *testing.T) {
	type fields struct {
		ctx          context.Context
		cancel       context.CancelFunc
		destinations []Destination
		location     *time.Location
		dateformat   string
		prefix       string
		logdebug     bool
		buffer       int
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
				ctx:          tt.fields.ctx,
				cancel:       tt.fields.cancel,
				destinations: tt.fields.destinations,
				location:     tt.fields.location,
				dateformat:   tt.fields.dateformat,
				prefix:       tt.fields.prefix,
				logdebug:     tt.fields.logdebug,
				buffer:       tt.fields.buffer,
			}
			l.Critf(tt.args.err, tt.args.format, tt.args.v...)
		})
	}
}

func Test_alog_Fatal(t *testing.T) {
	type fields struct {
		ctx          context.Context
		cancel       context.CancelFunc
		destinations []Destination
		location     *time.Location
		dateformat   string
		prefix       string
		logdebug     bool
		buffer       int
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
				ctx:          tt.fields.ctx,
				cancel:       tt.fields.cancel,
				destinations: tt.fields.destinations,
				location:     tt.fields.location,
				dateformat:   tt.fields.dateformat,
				prefix:       tt.fields.prefix,
				logdebug:     tt.fields.logdebug,
				buffer:       tt.fields.buffer,
			}
			l.Fatal(tt.args.err, tt.args.v...)
		})
	}
}

func Test_alog_Fatalln(t *testing.T) {
	type fields struct {
		ctx          context.Context
		cancel       context.CancelFunc
		destinations []Destination
		location     *time.Location
		dateformat   string
		prefix       string
		logdebug     bool
		buffer       int
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
				ctx:          tt.fields.ctx,
				cancel:       tt.fields.cancel,
				destinations: tt.fields.destinations,
				location:     tt.fields.location,
				dateformat:   tt.fields.dateformat,
				prefix:       tt.fields.prefix,
				logdebug:     tt.fields.logdebug,
				buffer:       tt.fields.buffer,
			}
			l.Fatalln(tt.args.err, tt.args.v...)
		})
	}
}

func Test_alog_Fatalf(t *testing.T) {
	type fields struct {
		ctx          context.Context
		cancel       context.CancelFunc
		destinations []Destination
		location     *time.Location
		dateformat   string
		prefix       string
		logdebug     bool
		buffer       int
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
				ctx:          tt.fields.ctx,
				cancel:       tt.fields.cancel,
				destinations: tt.fields.destinations,
				location:     tt.fields.location,
				dateformat:   tt.fields.dateformat,
				prefix:       tt.fields.prefix,
				logdebug:     tt.fields.logdebug,
				buffer:       tt.fields.buffer,
			}
			l.Fatalf(tt.args.err, tt.args.format, tt.args.v...)
		})
	}
}

func Test_alog_Custom(t *testing.T) {
	type fields struct {
		ctx          context.Context
		cancel       context.CancelFunc
		destinations []Destination
		location     *time.Location
		dateformat   string
		prefix       string
		logdebug     bool
		buffer       int
	}
	type args struct {
		ltype string
		err   error
		v     []interface{}
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
				ctx:          tt.fields.ctx,
				cancel:       tt.fields.cancel,
				destinations: tt.fields.destinations,
				location:     tt.fields.location,
				dateformat:   tt.fields.dateformat,
				prefix:       tt.fields.prefix,
				logdebug:     tt.fields.logdebug,
				buffer:       tt.fields.buffer,
			}
			l.Custom(tt.args.ltype, tt.args.err, tt.args.v...)
		})
	}
}

func Test_alog_Customln(t *testing.T) {
	type fields struct {
		ctx          context.Context
		cancel       context.CancelFunc
		destinations []Destination
		location     *time.Location
		dateformat   string
		prefix       string
		logdebug     bool
		buffer       int
	}
	type args struct {
		ltype string
		err   error
		v     []interface{}
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
				ctx:          tt.fields.ctx,
				cancel:       tt.fields.cancel,
				destinations: tt.fields.destinations,
				location:     tt.fields.location,
				dateformat:   tt.fields.dateformat,
				prefix:       tt.fields.prefix,
				logdebug:     tt.fields.logdebug,
				buffer:       tt.fields.buffer,
			}
			l.Customln(tt.args.ltype, tt.args.err, tt.args.v...)
		})
	}
}

func Test_alog_Customf(t *testing.T) {
	type fields struct {
		ctx          context.Context
		cancel       context.CancelFunc
		destinations []Destination
		location     *time.Location
		dateformat   string
		prefix       string
		logdebug     bool
		buffer       int
	}
	type args struct {
		ltype  string
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
				ctx:          tt.fields.ctx,
				cancel:       tt.fields.cancel,
				destinations: tt.fields.destinations,
				location:     tt.fields.location,
				dateformat:   tt.fields.dateformat,
				prefix:       tt.fields.prefix,
				logdebug:     tt.fields.logdebug,
				buffer:       tt.fields.buffer,
			}
			l.Customf(tt.args.ltype, tt.args.err, tt.args.format, tt.args.v...)
		})
	}
}

func Test_alog_AddOutput(t *testing.T) {
	type fields struct {
		ctx          context.Context
		cancel       context.CancelFunc
		destinations []Destination
		location     *time.Location
		dateformat   string
		prefix       string
		logdebug     bool
		buffer       int
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
				ctx:          tt.fields.ctx,
				cancel:       tt.fields.cancel,
				destinations: tt.fields.destinations,
				location:     tt.fields.location,
				dateformat:   tt.fields.dateformat,
				prefix:       tt.fields.prefix,
				logdebug:     tt.fields.logdebug,
				buffer:       tt.fields.buffer,
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
		ctx          context.Context
		cancel       context.CancelFunc
		destinations []Destination
		location     *time.Location
		dateformat   string
		prefix       string
		logdebug     bool
		buffer       int
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
				ctx:          tt.fields.ctx,
				cancel:       tt.fields.cancel,
				destinations: tt.fields.destinations,
				location:     tt.fields.location,
				dateformat:   tt.fields.dateformat,
				prefix:       tt.fields.prefix,
				logdebug:     tt.fields.logdebug,
				buffer:       tt.fields.buffer,
			}
			l.Close()
		})
	}
}

func Test_alog_Validate(t *testing.T) {
	type fields struct {
		ctx          context.Context
		cancel       context.CancelFunc
		destinations []Destination
		location     *time.Location
		dateformat   string
		prefix       string
		logdebug     bool
		buffer       int
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
				ctx:          tt.fields.ctx,
				cancel:       tt.fields.cancel,
				destinations: tt.fields.destinations,
				location:     tt.fields.location,
				dateformat:   tt.fields.dateformat,
				prefix:       tt.fields.prefix,
				logdebug:     tt.fields.logdebug,
				buffer:       tt.fields.buffer,
			}
			if gotValid := l.Validate(); gotValid != tt.wantValid {
				t.Errorf("alog.Validate() = %v, want %v", gotValid, tt.wantValid)
			}
		})
	}
}

func Test_alog_buildlog(t *testing.T) {
	type fields struct {
		ctx          context.Context
		cancel       context.CancelFunc
		destinations []Destination
		location     *time.Location
		dateformat   string
		prefix       string
		logdebug     bool
		buffer       int
	}
	type args struct {
		logtype int8
		custom  string
		err     error
		format  *string
		v       []interface{}
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantNewlog log
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &alog{
				ctx:          tt.fields.ctx,
				cancel:       tt.fields.cancel,
				destinations: tt.fields.destinations,
				location:     tt.fields.location,
				dateformat:   tt.fields.dateformat,
				prefix:       tt.fields.prefix,
				logdebug:     tt.fields.logdebug,
				buffer:       tt.fields.buffer,
			}
			if gotNewlog := l.buildlog(tt.args.logtype, tt.args.custom, tt.args.err, tt.args.format, tt.args.v...); !reflect.DeepEqual(gotNewlog, tt.wantNewlog) {
				t.Errorf("alog.buildlog() = %v, want %v", gotNewlog, tt.wantNewlog)
			}
		})
	}
}

func Test_alog_clog(t *testing.T) {
	type fields struct {
		ctx          context.Context
		cancel       context.CancelFunc
		destinations []Destination
		location     *time.Location
		dateformat   string
		prefix       string
		logdebug     bool
		buffer       int
	}
	type args struct {
		ctx     context.Context
		v       <-chan interface{}
		logtype int8
		custom  string
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
				ctx:          tt.fields.ctx,
				cancel:       tt.fields.cancel,
				destinations: tt.fields.destinations,
				location:     tt.fields.location,
				dateformat:   tt.fields.dateformat,
				prefix:       tt.fields.prefix,
				logdebug:     tt.fields.logdebug,
				buffer:       tt.fields.buffer,
			}
			l.clog(tt.args.ctx, tt.args.v, tt.args.logtype, tt.args.custom)
		})
	}
}

func Test_alog_Printc(t *testing.T) {
	type fields struct {
		ctx          context.Context
		cancel       context.CancelFunc
		destinations []Destination
		location     *time.Location
		dateformat   string
		prefix       string
		logdebug     bool
		buffer       int
	}
	type args struct {
		ctx context.Context
		v   <-chan interface{}
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
				ctx:          tt.fields.ctx,
				cancel:       tt.fields.cancel,
				destinations: tt.fields.destinations,
				location:     tt.fields.location,
				dateformat:   tt.fields.dateformat,
				prefix:       tt.fields.prefix,
				logdebug:     tt.fields.logdebug,
				buffer:       tt.fields.buffer,
			}
			l.Printc(tt.args.ctx, tt.args.v)
		})
	}
}

func Test_alog_Debugc(t *testing.T) {
	type fields struct {
		ctx          context.Context
		cancel       context.CancelFunc
		destinations []Destination
		location     *time.Location
		dateformat   string
		prefix       string
		logdebug     bool
		buffer       int
	}
	type args struct {
		ctx context.Context
		v   <-chan interface{}
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
				ctx:          tt.fields.ctx,
				cancel:       tt.fields.cancel,
				destinations: tt.fields.destinations,
				location:     tt.fields.location,
				dateformat:   tt.fields.dateformat,
				prefix:       tt.fields.prefix,
				logdebug:     tt.fields.logdebug,
				buffer:       tt.fields.buffer,
			}
			l.Debugc(tt.args.ctx, tt.args.v)
		})
	}
}

func Test_alog_Warnc(t *testing.T) {
	type fields struct {
		ctx          context.Context
		cancel       context.CancelFunc
		destinations []Destination
		location     *time.Location
		dateformat   string
		prefix       string
		logdebug     bool
		buffer       int
	}
	type args struct {
		ctx context.Context
		v   <-chan interface{}
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
				ctx:          tt.fields.ctx,
				cancel:       tt.fields.cancel,
				destinations: tt.fields.destinations,
				location:     tt.fields.location,
				dateformat:   tt.fields.dateformat,
				prefix:       tt.fields.prefix,
				logdebug:     tt.fields.logdebug,
				buffer:       tt.fields.buffer,
			}
			l.Warnc(tt.args.ctx, tt.args.v)
		})
	}
}

func Test_alog_Errorc(t *testing.T) {
	type fields struct {
		ctx          context.Context
		cancel       context.CancelFunc
		destinations []Destination
		location     *time.Location
		dateformat   string
		prefix       string
		logdebug     bool
		buffer       int
	}
	type args struct {
		ctx context.Context
		v   <-chan interface{}
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
				ctx:          tt.fields.ctx,
				cancel:       tt.fields.cancel,
				destinations: tt.fields.destinations,
				location:     tt.fields.location,
				dateformat:   tt.fields.dateformat,
				prefix:       tt.fields.prefix,
				logdebug:     tt.fields.logdebug,
				buffer:       tt.fields.buffer,
			}
			l.Errorc(tt.args.ctx, tt.args.v)
		})
	}
}

func Test_alog_Critc(t *testing.T) {
	type fields struct {
		ctx          context.Context
		cancel       context.CancelFunc
		destinations []Destination
		location     *time.Location
		dateformat   string
		prefix       string
		logdebug     bool
		buffer       int
	}
	type args struct {
		ctx context.Context
		v   <-chan interface{}
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
				ctx:          tt.fields.ctx,
				cancel:       tt.fields.cancel,
				destinations: tt.fields.destinations,
				location:     tt.fields.location,
				dateformat:   tt.fields.dateformat,
				prefix:       tt.fields.prefix,
				logdebug:     tt.fields.logdebug,
				buffer:       tt.fields.buffer,
			}
			l.Critc(tt.args.ctx, tt.args.v)
		})
	}
}

func Test_alog_Customc(t *testing.T) {
	type fields struct {
		ctx          context.Context
		cancel       context.CancelFunc
		destinations []Destination
		location     *time.Location
		dateformat   string
		prefix       string
		logdebug     bool
		buffer       int
	}
	type args struct {
		ctx   context.Context
		v     <-chan interface{}
		ltype string
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
				ctx:          tt.fields.ctx,
				cancel:       tt.fields.cancel,
				destinations: tt.fields.destinations,
				location:     tt.fields.location,
				dateformat:   tt.fields.dateformat,
				prefix:       tt.fields.prefix,
				logdebug:     tt.fields.logdebug,
				buffer:       tt.fields.buffer,
			}
			l.Customc(tt.args.ctx, tt.args.v, tt.args.ltype)
		})
	}
}
