package main

import (
	"context"
	"reflect"
	"testing"

	"github.com/EdwardShen125/bytedance-simple-server/kitex_gen/favorite"
)

func TestFavoriteServiceImpl_FavoriteAction(t *testing.T) {
	type args struct {
		ctx context.Context
		req *favorite.FavoriteActionRequest
	}
	tests := []struct {
		name     string
		args     args
		wantResp *favorite.FavoriteActionResponse
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &FavoriteServiceImpl{}
			gotResp, err := s.FavoriteAction(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("FavoriteAction() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("FavoriteAction() gotResp = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}

func TestFavoriteServiceImpl_FavoriteList(t *testing.T) {
	type args struct {
		ctx context.Context
		req *favorite.FavoriteListRequest
	}
	tests := []struct {
		name     string
		args     args
		wantResp *favorite.FavoriteListResponse
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &FavoriteServiceImpl{}
			gotResp, err := s.FavoriteList(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("FavoriteList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("FavoriteList() gotResp = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}
