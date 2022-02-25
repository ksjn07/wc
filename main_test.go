package main

import (
	"bytes"
	"io"
	"testing"
)

func TestCounter(t *testing.T) {
	type args struct {
		r io.Reader
		f bool
		b string
	}
	a1 := args{bytes.NewBufferString("this is a test for word count"), false, ""}
	a2 := args{bytes.NewBufferString("this is a\n test for\n word count"), true, ""}
	a3 := args{bytes.NewBufferString("this is a {} test for\n word count"), true, "{"}
	tests := []struct {
		name  string
		args  args
		want1 int
		want2 int
	}{
		{"test when a normal string is provided", a1, 0, 7},
		{"test when lines are present", a2, 0, 3},
		{"test when lines are present", a3, 1, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got1, got2 := Counter(tt.args.r, tt.args.f, tt.args.b); got1 != tt.want1 || got2 != tt.want2 {
				t.Errorf("Counter() = %v, want %v Counter()=%v, want %v", got1, tt.want1, got2, tt.want2)
			}
		})
	}
}
