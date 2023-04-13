package main

import (
	"context"
	"reflect"
	"testing"

	"github.com/EdwardShen125/bytedance-simple-server/kitex_gen/comment"
)

func TestCommentServiceImpl_CommentList(t *testing.T) {
	type args struct {
		ctx context.Context
		req *comment.CommentListRequest
	}
	tests := []struct {
		name     string
		args     args
		wantResp *comment.CommentListResponse
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &CommentServiceImpl{}
			gotResp, err := s.CommentList(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("CommentList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("CommentList() gotResp = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}

func TestCommentServiceImpl_CommentAction(t *testing.T) {
	type args struct {
		ctx context.Context
		req *comment.CommentActionRequest
	}
	tests := []struct {
		name     string
		args     args
		wantResp *comment.CommentActionResponse
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &CommentServiceImpl{}
			gotResp, err := s.CommentAction(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("CommentAction() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("CommentAction() gotResp = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}