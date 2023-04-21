package main

import (
	"context"
	"reflect"
	"testing"

	"github.com/EdwardShen125/bytedance-simple-server/kitex_gen/publish"
)

func TestPublishServiceImpl_PublishAction(t *testing.T) {
	type args struct {
		ctx context.Context
		req *publish.PublishActionRequest
	}
	tests := []struct {
		name     string
		args     args
		wantResp *publish.PublishActionResponse
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &PublishServiceImpl{}
			gotResp, err := s.PublishAction(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("PublishAction() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("PublishAction() gotResp = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}

func TestPublishServiceImpl_PublishList(t *testing.T) {
	type args struct {
		ctx context.Context
		req *publish.PublishListRequest
	}
	tests := []struct {
		name     string
		args     args
		wantResp *publish.PublishListResponse
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &PublishServiceImpl{}
			gotResp, err := s.PublishList(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("PublishList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("PublishList() gotResp = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}
