package tracer

import "testing"

func TestInitJaeger(t *testing.T) {
	type args struct {
		service string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			InitJaeger(tt.args.service)
		})
	}
}
