package middleware

import (
	"github.com/cloudwego/kitex/pkg/endpoint"
	"reflect"
	"testing"
)

func TestClientMiddleware(t *testing.T) {
	type args struct {
		next endpoint.Endpoint
	}
	tests := []struct {
		name string
		args args
		want endpoint.Endpoint
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ClientMiddleware(tt.args.next); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ClientMiddleware() = %v, want %v", got, tt.want)
			}
		})
	}
}
