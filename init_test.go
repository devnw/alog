package alog

import (
	"context"
	"io"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		ctx    context.Context
		format int
		prefix string
		out    []io.Writer
	}
	tests := []struct {
		name    string
		args    args
		want    Logger
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.args.ctx, tt.args.format, tt.args.prefix, tt.args.out...)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
