package alog

import (
	"context"
	"testing"
)

func TestPrintc(t *testing.T) {
	type args struct {
		ctx context.Context
		v   <-chan interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Printc(tt.args.ctx, tt.args.v)
		})
	}
}

func TestPrint(t *testing.T) {
	type args struct {
		v []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Print(tt.args.v...)
		})
	}
}

func TestPrintln(t *testing.T) {
	type args struct {
		v []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Println(tt.args.v...)
		})
	}
}

func TestPrintf(t *testing.T) {
	type args struct {
		format string
		v      []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Printf(tt.args.format, tt.args.v...)
		})
	}
}

func TestDebugc(t *testing.T) {
	type args struct {
		ctx context.Context
		v   <-chan interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Debugc(tt.args.ctx, tt.args.v)
		})
	}
}

func TestDebug(t *testing.T) {
	type args struct {
		v []interface{}
	}
	tests := []struct {
		name string
		err  error
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Debug(tt.err, tt.args.v...)
		})
	}
}

func TestDebugln(t *testing.T) {
	type args struct {
		v []interface{}
	}
	tests := []struct {
		name string
		err  error
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Debugln(tt.err, tt.args.v...)
		})
	}
}

func TestDebugf(t *testing.T) {
	type args struct {
		format string
		v      []interface{}
	}
	tests := []struct {
		name string
		err  error
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Debugf(tt.err, tt.args.format, tt.args.v...)
		})
	}
}

func TestWarnc(t *testing.T) {
	type args struct {
		ctx context.Context
		v   <-chan interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Warnc(tt.args.ctx, tt.args.v)
		})
	}
}

func TestWarn(t *testing.T) {
	type args struct {
		err error
		v   []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Warn(tt.args.err, tt.args.v...)
		})
	}
}

func TestWarnln(t *testing.T) {
	type args struct {
		err error
		v   []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Warnln(tt.args.err, tt.args.v...)
		})
	}
}

func TestWarnf(t *testing.T) {
	type args struct {
		err    error
		format string
		v      []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Warnf(tt.args.err, tt.args.format, tt.args.v...)
		})
	}
}

func TestErrorc(t *testing.T) {
	type args struct {
		ctx context.Context
		v   <-chan interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Errorc(tt.args.ctx, tt.args.v)
		})
	}
}

func TestError(t *testing.T) {
	type args struct {
		err error
		v   []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Error(tt.args.err, tt.args.v...)
		})
	}
}

func TestErrorln(t *testing.T) {
	type args struct {
		err error
		v   []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Errorln(tt.args.err, tt.args.v...)
		})
	}
}

func TestErrorf(t *testing.T) {
	type args struct {
		err    error
		format string
		v      []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Errorf(tt.args.err, tt.args.format, tt.args.v...)
		})
	}
}

func TestCritc(t *testing.T) {
	type args struct {
		ctx context.Context
		v   <-chan interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Critc(tt.args.ctx, tt.args.v)
		})
	}
}

func TestCrit(t *testing.T) {
	type args struct {
		err error
		v   []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Crit(tt.args.err, tt.args.v...)
		})
	}
}

func TestCritln(t *testing.T) {
	type args struct {
		err error
		v   []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Critln(tt.args.err, tt.args.v...)
		})
	}
}

func TestCritf(t *testing.T) {
	type args struct {
		err    error
		format string
		v      []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Critf(tt.args.err, tt.args.format, tt.args.v...)
		})
	}
}

func TestFatal(t *testing.T) {
	type args struct {
		err error
		v   []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Fatal(tt.args.err, tt.args.v...)
		})
	}
}

func TestFatalln(t *testing.T) {
	type args struct {
		err error
		v   []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Fatalln(tt.args.err, tt.args.v...)
		})
	}
}

func TestFatalf(t *testing.T) {
	type args struct {
		err    error
		format string
		v      []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Fatalf(tt.args.err, tt.args.format, tt.args.v...)
		})
	}
}

func TestCustomc(t *testing.T) {
	type args struct {
		ctx   context.Context
		v     <-chan interface{}
		ltype string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Customc(tt.args.ctx, tt.args.v, tt.args.ltype)
		})
	}
}

func TestCustom(t *testing.T) {
	type args struct {
		ltype string
		err   error
		v     []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Custom(tt.args.ltype, tt.args.err, tt.args.v...)
		})
	}
}

func TestCustomln(t *testing.T) {
	type args struct {
		ltype string
		err   error
		v     []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Customln(tt.args.ltype, tt.args.err, tt.args.v...)
		})
	}
}

func TestCustomf(t *testing.T) {
	type args struct {
		ltype  string
		err    error
		format string
		v      []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Customf(tt.args.ltype, tt.args.err, tt.args.format, tt.args.v...)
		})
	}
}
