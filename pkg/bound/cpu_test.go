package bound

import (
	"context"
	"github.com/cloudwego/kitex/pkg/remote"
	"net"
	"reflect"
	"testing"
)

func TestNewCpuLimitHandler(t *testing.T) {
	tests := []struct {
		name string
		want remote.InboundHandler
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCpuLimitHandler(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCpuLimitHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_cpuLimitHandler_OnActive(t *testing.T) {
	type args struct {
		ctx  context.Context
		conn net.Conn
	}
	tests := []struct {
		name    string
		args    args
		want    context.Context
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &cpuLimitHandler{}
			got, err := c.OnActive(tt.args.ctx, tt.args.conn)
			if (err != nil) != tt.wantErr {
				t.Errorf("OnActive() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OnActive() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_cpuLimitHandler_OnInactive(t *testing.T) {
	type args struct {
		ctx  context.Context
		conn net.Conn
	}
	tests := []struct {
		name string
		args args
		want context.Context
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &cpuLimitHandler{}
			if got := c.OnInactive(tt.args.ctx, tt.args.conn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OnInactive() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_cpuLimitHandler_OnMessage(t *testing.T) {
	type args struct {
		ctx    context.Context
		args   remote.Message
		result remote.Message
	}
	tests := []struct {
		name    string
		args    args
		want    context.Context
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &cpuLimitHandler{}
			got, err := c.OnMessage(tt.args.ctx, tt.args.args, tt.args.result)
			if (err != nil) != tt.wantErr {
				t.Errorf("OnMessage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OnMessage() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_cpuLimitHandler_OnRead(t *testing.T) {
	type args struct {
		ctx  context.Context
		conn net.Conn
	}
	tests := []struct {
		name    string
		args    args
		want    context.Context
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &cpuLimitHandler{}
			got, err := c.OnRead(tt.args.ctx, tt.args.conn)
			if (err != nil) != tt.wantErr {
				t.Errorf("OnRead() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OnRead() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_cpuPercent(t *testing.T) {
	tests := []struct {
		name string
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cpuPercent(); got != tt.want {
				t.Errorf("cpuPercent() = %v, want %v", got, tt.want)
			}
		})
	}
}
