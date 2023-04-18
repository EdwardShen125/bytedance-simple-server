package main

import (
	"context"
	"reflect"
	"testing"

	"github.com/EdwardShen125/bytedance-simple-server/kitex_gen/feed"
)

func TestFeedServiceImpl_Feed(t *testing.T) {
	type args struct {
		ctx context.Context
		req *feed.FeedRequest
	}
	tests := []struct {
		name     string
		args     args
		wantResp *feed.FeedResponse
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &FeedServiceImpl{}
			gotResp, err := s.Feed(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Feed() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("Feed() gotResp = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}
