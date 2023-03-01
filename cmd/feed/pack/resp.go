package pack

import (
	"errors"
	"github.com/EdwardShen125/bytedance-simple-server/kitex_gen/feed"
	"github.com/EdwardShen125/bytedance-simple-server/pkg/errno"
)

func feedResp(err errno.ErrNo) *feed.FeedResponse {
	return &feed.FeedResponse{StatusCode: err.ErrCode, StatusMsg: &err.ErrMsg}
}

func BuildFeedResp(err error) *feed.FeedResponse {
	if err == nil {
		return feedResp(errno.Success)
	}
	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return feedResp(e)
	}
	s := errno.ServiceErr.WithMessage(err.Error())
	return feedResp(s)
}
