package main

import (
	"context"
	"unicode/utf8"

	"github.com/EdwardShen125/bytedance-simple-server/cmd/comment/pack"
	"github.com/EdwardShen125/bytedance-simple-server/cmd/comment/service"
	"github.com/EdwardShen125/bytedance-simple-server/kitex_gen/comment"
	"github.com/EdwardShen125/bytedance-simple-server/pkg/errno"
)

type CommentServiceImpl struct{}

func (s *CommentServiceImpl) CommentList(ctx context.Context, req *comment.CommentListRequest) (resp *comment.CommentListResponse, err error) {
	resp = new(comment.CommentListResponse)

	if req.VideoId == 0 {
		resp = pack.BuildCommentListResp(errno.ParamErr)
		return resp, nil
	}

	commentList, err := service.NewCommentListService(ctx).CommentList(req)
	if err != nil {
		resp = pack.BuildCommentListResp(err)
		return resp, nil
	}

	resp = pack.BuildCommentListResp(errno.Success)
	resp.CommentList = commentList
	return resp, nil
}

func (s *CommentServiceImpl) CommentAction(ctx context.Context, req *comment.CommentActionRequest) (resp *comment.CommentActionResponse, err error) {
	resp = new(comment.CommentActionResponse)

	if len(req.Token) == 0 || req.VideoId == 0 {
		resp = pack.BuildCommentActionResp(errno.ParamErr)
		return resp, nil
	}

	if req.CommentId != nil && *req.CommentId == 0 {
		resp = pack.BuildCommentActionResp(errno.ParamErr)
		return resp, nil
	}

	if req.CommentText != nil && utf8.RuneCountInString(*req.CommentText) > 20 {
		resp = pack.BuildCommentActionResp(errno.ParamErr)
		return resp, nil
	}

	comment_, err := service.NewCommentActionService(ctx).CommentAction(req)
	if err != nil {
		resp = pack.BuildCommentActionResp(err)
		return resp, nil
	}

	resp = pack.BuildCommentActionResp(errno.Success)
	resp.Comment = comment_
	return resp, nil
}
