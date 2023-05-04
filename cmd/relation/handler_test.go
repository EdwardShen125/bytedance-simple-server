package main

import (
	"context"
	"reflect"
	"testing"

	"github.com/EdwardShen125/bytedance-simple-server/kitex_gen/relation"
)

func TestRelationServiceImpl_RelationAction(t *testing.T) {
	type args struct {
		ctx context.Context
		req *relation.RelationActionRequest
	}
	tests := []struct {
		name     string
		args     args
		wantResp *relation.RelationActionResponse
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &RelationServiceImpl{}
			gotResp, err := s.RelationAction(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("RelationAction() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("RelationAction() gotResp = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}

func TestRelationServiceImpl_RelationFollowList(t *testing.T) {
	type args struct {
		ctx context.Context
		req *relation.RelationFollowListRequest
	}
	tests := []struct {
		name     string
		args     args
		wantResp *relation.RelationFollowListResponse
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &RelationServiceImpl{}
			gotResp, err := s.RelationFollowList(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("RelationFollowList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("RelationFollowList() gotResp = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}

func TestRelationServiceImpl_RelationFollowerList(t *testing.T) {
	type args struct {
		ctx context.Context
		req *relation.RelationFollowerListRequest
	}
	tests := []struct {
		name     string
		args     args
		wantResp *relation.RelationFollowerListResponse
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &RelationServiceImpl{}
			gotResp, err := s.RelationFollowerList(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("RelationFollowerList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("RelationFollowerList() gotResp = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}

func TestRelationServiceImpl_RelationFriendList(t *testing.T) {
	type args struct {
		ctx context.Context
		req *relation.RelationFriendListRequest
	}
	tests := []struct {
		name     string
		args     args
		wantResp *relation.RelationFriendListResponse
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &RelationServiceImpl{}
			gotResp, err := s.RelationFriendList(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("RelationFriendList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("RelationFriendList() gotResp = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}
