package pack

import (
	"errors"
	"github.com/EdwardShen125/bytedance-simple-server/kitex_gen/comment"
	"github.com/EdwardShen125/bytedance-simple-server/pkg/errno"
)

func BuildCommentListResp(err error) *comment.CommentListResponse {
	if err == nil {
		return commentListResp(errno.Success)
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return commentListResp(e)
	}

	s := errno.ServiceErr.WithMessage(err.Error())
	return commentListResp(s)
}

func commentListResp(err errno.ErrNo) *comment.CommentListResponse {
	return &comment.CommentListResponse{StatusCode: err.ErrCode, StatusMsg: &err.ErrMsg}
}

func BuildCommentActionResp(err error) *comment.CommentActionResponse {
	if err == nil {
		return commentActionResp(errno.Success)
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return commentActionResp(e)
	}

	s := errno.ServiceErr.WithMessage(err.Error())
	return commentActionResp(s)
}

func commentActionResp(err errno.ErrNo) *comment.CommentActionResponse {
	return &comment.CommentActionResponse{StatusCode: err.ErrCode, StatusMsg: &err.ErrMsg}
}

