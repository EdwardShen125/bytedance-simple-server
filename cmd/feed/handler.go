package main

import (
	"context"
	"github.com/EdwardShen125/bytedance-simple-server/cmd/feed/pack"
	"github.com/EdwardShen125/bytedance-simple-server/cmd/feed/service"
	"github.com/EdwardShen125/bytedance-simple-server/kitex_gen/feed"
	"github.com/EdwardShen125/bytedance-simple-server/pkg/errno"
)

type FeedServiceImpl struct{}

// Feed implements the FeedServiceImpl interface.
func (s *FeedServiceImpl) Feed(ctx context.Context, req *feed.FeedRequest) (resp *feed.FeedResponse, err error) {
	resp = new(feed.FeedResponse)

	if *req.LatestTime <= 0 {
		resp = pack.BuildFeedResp(errno.ParamErr)
		return resp, nil
	}

	videos, nextTime, err := service.NewFeedService(ctx).Feed(req)
	if err != nil {
		resp = pack.BuildFeedResp(err)
		return resp, nil
	}
	resp = pack.BuildFeedResp(errno.Success)
	resp.VideoList = videos
	resp.NextTime = &nextTime
	return resp, nil
}

