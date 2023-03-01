package main

import (
	"context"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/EdwardShen125/bytedance-simple-server/cmd/user/pack"
	"github.com/EdwardShen125/bytedance-simple-server/cmd/user/service"
	"github.com/EdwardShen125/bytedance-simple-server/kitex_gen/user"
	"github.com/EdwardShen125/bytedance-simple-server/pkg/errno"
	"github.com/EdwardShen125/bytedance-simple-server/pkg/jwt"
)

type UserServiceImpl struct{}

func (s *UserServiceImpl) UserLogin(ctx context.Context, req *user.UserLoginRequest) (resp *user.UserLoginResponse, err error) {
	resp = new(user.UserLoginResponse)

	if len(req.Username) == 0 || len(req.Password) == 0 {
		resp = pack.BuildUserLoginResp(errno.ParamErr)
		return resp, nil
	}

	uid, err := service.NewLoginUserService(ctx).CheckUser(req)
	if err != nil {
		resp = pack.BuildUserLoginResp(err)
		return resp, nil
	}
	resp = pack.BuildUserLoginResp(errno.Success)
	resp.UserId = uid
	return resp, nil
}

func (s *UserServiceImpl) UserRegister(ctx context.Context, req *user.UserRegisterRequest) (resp *user.UserRegisterResponse, err error) {
	resp = new(user.UserRegisterResponse)

	if len(req.Username) == 0 || len(req.Password) == 0 {
		resp = pack.BuildUserRegisterResp(errno.ParamErr)
		return
	}

	userId, err := service.NewRegisterUserService(ctx).RegisterUser(req)
	if err != nil {
		resp = pack.BuildUserRegisterResp(err)
		return
	}

	token, err := Jwt.CreateToken(jwt.CustomClaims{
		Id: userId,
	})
	if err != nil {
		resp = pack.BuildUserRegisterResp(err)
		return
	}

	resp = pack.BuildUserRegisterResp(errno.Success)
	resp.UserId = userId
	resp.Token = token
	return resp, nil
}

func (s *UserServiceImpl) UserInfo(ctx context.Context, req *user.UserRequest) (resp *user.UserResponse, err error) {
	resp = new(user.UserResponse)

	klog.CtxInfof(ctx, "UserInfo req %+v\n", req)

	if req.UserId == 0 {
		resp = pack.BuildUserInfoResp(errno.ParamErr)
		return
	}

	user_, err := service.NewUserInfoService(ctx).UserInfo(req)
	if err != nil {
		resp = pack.BuildUserInfoResp(err)
		return
	}

	resp = pack.BuildUserInfoResp(errno.Success)
	resp.User = user_
	return resp, nil
}
