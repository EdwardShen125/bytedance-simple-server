package main

import (
	"context"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/EdwardShen125/bytedance-simple-server/cmd/favorite/pack"
	"github.com/EdwardShen125/bytedance-simple-server/cmd/favorite/service"
	"github.com/EdwardShen125/bytedance-simple-server/kitex_gen/favorite"
	"github.com/EdwardShen125/bytedance-simple-server/pkg/errno"
)

type FavoriteServiceImpl struct{}

// FavoriteAction implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) FavoriteAction(ctx context.Context, req *favorite.FavoriteActionRequest) (resp *favorite.FavoriteActionResponse, err error) {
	resp = new(favorite.FavoriteActionResponse)

	if len(req.Token) == 0 || req.VideoId == 0 || req.ActionType == 0 {
		resp = pack.BuildFavoriteActionResp(errno.ParamErr)
		return resp, nil
	}

	err = service.NewFavoriteActionService(ctx).FavoriteAction(req)
	if err != nil {
		resp = pack.BuildFavoriteActionResp(err)
		return resp, nil
	}
	resp = pack.BuildFavoriteActionResp(errno.Success)
	return resp, nil
}

// FavoriteList implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) FavoriteList(ctx context.Context, req *favorite.FavoriteListRequest) (resp *favorite.FavoriteListResponse, err error) {
	resp = new(favorite.FavoriteListResponse)

	klog.Infof("FavoriteList req %+v\n", req)
	if req.UserId == 0 {
		resp = pack.BuildFavoriteListResp(errno.ParamErr)
		return resp, nil
	}

	videoList, err := service.NewFavoriteListService(ctx).FavoriteList(req)
	if err != nil {
		resp = pack.BuildFavoriteListResp(err)
		return resp, nil
	}

	resp = pack.BuildFavoriteListResp(errno.Success)
	resp.VideoList = videoList
	return resp, nil
}
