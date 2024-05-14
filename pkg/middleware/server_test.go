package middleware

import (
	"github.com/cloudwego/kitex/pkg/endpoint"
	"reflect"
	"testing"
)

func TestServerMiddleware(t *testing.T) {
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
			if got := ServerMiddleware(tt.args.next); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ServerMiddleware() = %v, want %v", got, tt.want)
			}
		})
	}
}
