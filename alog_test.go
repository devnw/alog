package alog

import (
	"bytes"
	"context"
	"io"
	"testing"
)

func Test_alog_Print(t *testing.T) {
	type fields struct {
		ctx     context.Context
		cancel  context.CancelFunc
		outputs []io.Writer
		format  int
		prefix  string
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
				ctx:     tt.fields.ctx,
				cancel:  tt.fields.cancel,
				outputs: tt.fields.outputs,
				format:  tt.fields.format,
				prefix:  tt.fields.prefix,
			}
			l.Print(tt.args.v...)
		})
	}
}

func Test_alog_Println(t *testing.T) {
	type fields struct {
		ctx     context.Context
		cancel  context.CancelFunc
		outputs []io.Writer
		format  int
		prefix  string
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
				ctx:     tt.fields.ctx,
				cancel:  tt.fields.cancel,
				outputs: tt.fields.outputs,
				format:  tt.fields.format,
				prefix:  tt.fields.prefix,
			}
			l.Println(tt.args.v...)
		})
	}
}

func Test_alog_Printf(t *testing.T) {
	type fields struct {
		ctx     context.Context
		cancel  context.CancelFunc
		outputs []io.Writer
		format  int
		prefix  string
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
				ctx:     tt.fields.ctx,
				cancel:  tt.fields.cancel,
				outputs: tt.fields.outputs,
				format:  tt.fields.format,
				prefix:  tt.fields.prefix,
			}
			l.Printf(tt.args.format, tt.args.v...)
		})
	}
}

func Test_alog_Warn(t *testing.T) {
	type fields struct {
		ctx     context.Context
		cancel  context.CancelFunc
		outputs []io.Writer
		format  int
		prefix  string
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
				ctx:     tt.fields.ctx,
				cancel:  tt.fields.cancel,
				outputs: tt.fields.outputs,
				format:  tt.fields.format,
				prefix:  tt.fields.prefix,
			}
			l.Warn(tt.args.err, tt.args.v...)
		})
	}
}

func Test_alog_Warnln(t *testing.T) {
	type fields struct {
		ctx     context.Context
		cancel  context.CancelFunc
		outputs []io.Writer
		format  int
		prefix  string
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
				ctx:     tt.fields.ctx,
				cancel:  tt.fields.cancel,
				outputs: tt.fields.outputs,
				format:  tt.fields.format,
				prefix:  tt.fields.prefix,
			}
			l.Warnln(tt.args.err, tt.args.v...)
		})
	}
}

func Test_alog_Warnf(t *testing.T) {
	type fields struct {
		ctx     context.Context
		cancel  context.CancelFunc
		outputs []io.Writer
		format  int
		prefix  string
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
				ctx:     tt.fields.ctx,
				cancel:  tt.fields.cancel,
				outputs: tt.fields.outputs,
				format:  tt.fields.format,
				prefix:  tt.fields.prefix,
			}
			l.Warnf(tt.args.err, tt.args.format, tt.args.v...)
		})
	}
}

func Test_alog_Error(t *testing.T) {
	type fields struct {
		ctx     context.Context
		cancel  context.CancelFunc
		outputs []io.Writer
		format  int
		prefix  string
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
				ctx:     tt.fields.ctx,
				cancel:  tt.fields.cancel,
				outputs: tt.fields.outputs,
				format:  tt.fields.format,
				prefix:  tt.fields.prefix,
			}
			l.Error(tt.args.err, tt.args.v...)
		})
	}
}

func Test_alog_Errorln(t *testing.T) {
	type fields struct {
		ctx     context.Context
		cancel  context.CancelFunc
		outputs []io.Writer
		format  int
		prefix  string
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
				ctx:     tt.fields.ctx,
				cancel:  tt.fields.cancel,
				outputs: tt.fields.outputs,
				format:  tt.fields.format,
				prefix:  tt.fields.prefix,
			}
			l.Errorln(tt.args.err, tt.args.v...)
		})
	}
}

func Test_alog_Errorf(t *testing.T) {
	type fields struct {
		ctx     context.Context
		cancel  context.CancelFunc
		outputs []io.Writer
		format  int
		prefix  string
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
				ctx:     tt.fields.ctx,
				cancel:  tt.fields.cancel,
				outputs: tt.fields.outputs,
				format:  tt.fields.format,
				prefix:  tt.fields.prefix,
			}
			l.Errorf(tt.args.err, tt.args.format, tt.args.v...)
		})
	}
}

func Test_alog_Crit(t *testing.T) {
	type fields struct {
		ctx     context.Context
		cancel  context.CancelFunc
		outputs []io.Writer
		format  int
		prefix  string
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
				ctx:     tt.fields.ctx,
				cancel:  tt.fields.cancel,
				outputs: tt.fields.outputs,
				format:  tt.fields.format,
				prefix:  tt.fields.prefix,
			}
			l.Crit(tt.args.err, tt.args.v...)
		})
	}
}

func Test_alog_Critln(t *testing.T) {
	type fields struct {
		ctx     context.Context
		cancel  context.CancelFunc
		outputs []io.Writer
		format  int
		prefix  string
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
				ctx:     tt.fields.ctx,
				cancel:  tt.fields.cancel,
				outputs: tt.fields.outputs,
				format:  tt.fields.format,
				prefix:  tt.fields.prefix,
			}
			l.Critln(tt.args.err, tt.args.v...)
		})
	}
}

func Test_alog_Critf(t *testing.T) {
	type fields struct {
		ctx     context.Context
		cancel  context.CancelFunc
		outputs []io.Writer
		format  int
		prefix  string
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
				ctx:     tt.fields.ctx,
				cancel:  tt.fields.cancel,
				outputs: tt.fields.outputs,
				format:  tt.fields.format,
				prefix:  tt.fields.prefix,
			}
			l.Critf(tt.args.err, tt.args.format, tt.args.v...)
		})
	}
}

func Test_alog_Fatal(t *testing.T) {
	type fields struct {
		ctx     context.Context
		cancel  context.CancelFunc
		outputs []io.Writer
		format  int
		prefix  string
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
				ctx:     tt.fields.ctx,
				cancel:  tt.fields.cancel,
				outputs: tt.fields.outputs,
				format:  tt.fields.format,
				prefix:  tt.fields.prefix,
			}
			l.Fatal(tt.args.err, tt.args.v...)
		})
	}
}

func Test_alog_Fatalln(t *testing.T) {
	type fields struct {
		ctx     context.Context
		cancel  context.CancelFunc
		outputs []io.Writer
		format  int
		prefix  string
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
				ctx:     tt.fields.ctx,
				cancel:  tt.fields.cancel,
				outputs: tt.fields.outputs,
				format:  tt.fields.format,
				prefix:  tt.fields.prefix,
			}
			l.Fatalln(tt.args.err, tt.args.v...)
		})
	}
}

func Test_alog_Fatalf(t *testing.T) {
	type fields struct {
		ctx     context.Context
		cancel  context.CancelFunc
		outputs []io.Writer
		format  int
		prefix  string
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
				ctx:     tt.fields.ctx,
				cancel:  tt.fields.cancel,
				outputs: tt.fields.outputs,
				format:  tt.fields.format,
				prefix:  tt.fields.prefix,
			}
			l.Fatalf(tt.args.err, tt.args.format, tt.args.v...)
		})
	}
}

func Test_alog_AddOutput(t *testing.T) {
	type fields struct {
		ctx     context.Context
		cancel  context.CancelFunc
		outputs []io.Writer
		format  int
		prefix  string
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
				ctx:     tt.fields.ctx,
				cancel:  tt.fields.cancel,
				outputs: tt.fields.outputs,
				format:  tt.fields.format,
				prefix:  tt.fields.prefix,
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
		ctx     context.Context
		cancel  context.CancelFunc
		outputs []io.Writer
		format  int
		prefix  string
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
				ctx:     tt.fields.ctx,
				cancel:  tt.fields.cancel,
				outputs: tt.fields.outputs,
				format:  tt.fields.format,
				prefix:  tt.fields.prefix,
			}
			l.Close()
		})
	}
}

func Test_alog_Validate(t *testing.T) {
	type fields struct {
		ctx     context.Context
		cancel  context.CancelFunc
		outputs []io.Writer
		format  int
		prefix  string
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
				ctx:     tt.fields.ctx,
				cancel:  tt.fields.cancel,
				outputs: tt.fields.outputs,
				format:  tt.fields.format,
				prefix:  tt.fields.prefix,
			}
			if gotValid := l.Validate(); gotValid != tt.wantValid {
				t.Errorf("alog.Validate() = %v, want %v", gotValid, tt.wantValid)
			}
		})
	}
}
