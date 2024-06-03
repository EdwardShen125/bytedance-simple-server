package rpc

import (
	"context"
	"github.com/EdwardShen125/bytedance-simple-server/kitex_gen/comment"
	"reflect"
	"testing"
)

func TestCommentList(t *testing.T) {
	type args struct {
		ctx context.Context
		req *comment.CommentListRequest
	}
	tests := []struct {
		name    string
		args    args
		want    []*comment.Comment
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CommentList(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("CommentList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CommentList() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCommentAction(t *testing.T) {
	type args struct {
		ctx context.Context
		req *comment.CommentActionRequest
	}
	tests := []struct {
		name    string
		args    args
		want    *comment.Comment
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CommentAction(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("CommentAction() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CommentAction() got = %v, want %v", got, tt.want)
			}
		})
	}
}
