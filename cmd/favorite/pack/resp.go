package pack

import (
	"errors"
	"github.com/EdwardShen125/bytedance-simple-server/kitex_gen/favorite"
	"github.com/EdwardShen125/bytedance-simple-server/pkg/errno"
)

func BuildFavoriteActionResp(err error) *favorite.FavoriteActionResponse {
	if err == nil {
		return favoriteActionResp(errno.Success)
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return favoriteActionResp(e)
	}

	s := errno.ServiceErr.WithMessage(err.Error())
	return favoriteActionResp(s)
}

func favoriteActionResp(err errno.ErrNo) *favorite.FavoriteActionResponse {
	return &favorite.FavoriteActionResponse{StatusCode: err.ErrCode, StatusMsg: &err.ErrMsg}
}

func BuildFavoriteListResp(err error) *favorite.FavoriteListResponse {
	if err == nil {
		return favoriteListResp(errno.Success)
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return favoriteListResp(e)
	}

	s := errno.ServiceErr.WithMessage(err.Error())
	return favoriteListResp(s)
}

func favoriteListResp(err errno.ErrNo) *favorite.FavoriteListResponse {
	return &favorite.FavoriteListResponse{StatusCode: err.ErrCode, StatusMsg: &err.ErrMsg}
}