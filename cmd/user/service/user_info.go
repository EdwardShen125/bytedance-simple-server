package service

import (
	"context"
	"errors"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/EdwardShen125/bytedance-simple-server/cmd/user/dal/db"
	"github.com/EdwardShen125/bytedance-simple-server/cmd/user/pack"
	"github.com/EdwardShen125/bytedance-simple-server/kitex_gen/user"
	"github.com/EdwardShen125/bytedance-simple-server/pkg/constants"
	"github.com/EdwardShen125/bytedance-simple-server/pkg/jwt"
)

type UserInfoService struct {
	ctx context.Context
}

func NewUserInfoService(ctx context.Context) *UserInfoService {
	return &UserInfoService{
		ctx: ctx,
	}
}

func (s *UserInfoService) UserInfo(req *user.UserRequest) (*user.User, error) {
	Jwt := jwt.NewJWT([]byte(constants.SecretKey))
	currentId, err := Jwt.CheckToken(req.Token)
	if err != nil {
		return nil, err
	}

	userIds := []int64{req.UserId}
	klog.CtxInfof(context.TODO(), "UserInfo got req %+v. currentId %v", req, currentId)
	users, err := db.QueryUserByIds(s.ctx, userIds)
	if err != nil {
		return nil, err
	}
	if len(users) == 0 {
		return nil, errors.New("user not exist")
	}
	user_ := users[0]
	relationMap, err := db.QueryRelationByIds(s.ctx, currentId, userIds)
	if err != nil {
		return nil, err
	}

	var isFollow bool
	_, ok := relationMap[req.UserId]
	if ok {
		isFollow = true
	} else {
		isFollow = false
	}

	userInfo := pack.UserInfo(user_, isFollow)
	return userInfo, nil
}
