package alog

import (
	"reflect"
	"testing"
)

func TestStandards(t *testing.T) {
	tests := []struct {
		name string
		want []Destination
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Standards(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Standards() = %v, want %v", got, tt.want)
			}
		})
	}
}
