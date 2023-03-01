package pack

import (
	"errors"
	"github.com/EdwardShen125/bytedance-simple-server/kitex_gen/publish"
	"github.com/EdwardShen125/bytedance-simple-server/pkg/errno"
)

func BuildPublishActionResp(err error) *publish.PublishActionResponse {
	if err == nil {
		return publishActionResp(errno.Success)
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return publishActionResp(e)
	}

	s := errno.ServiceErr.WithMessage(err.Error())
	return publishActionResp(s)
}

func publishActionResp(err errno.ErrNo) *publish.PublishActionResponse {
	return &publish.PublishActionResponse{StatusCode: err.ErrCode, StatusMsg: &err.ErrMsg}
}

func BuildPublishListResp(err error) *publish.PublishListResponse {
	if err == nil {
		return publishListResp(errno.Success)
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return publishListResp(e)
	}

	s := errno.ServiceErr.WithMessage(err.Error())
	return publishListResp(s)
}

func publishListResp(err errno.ErrNo) *publish.PublishListResponse {
	return &publish.PublishListResponse{StatusCode: err.ErrCode, StatusMsg: &err.ErrMsg}
}