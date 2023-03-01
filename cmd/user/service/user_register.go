package service

import (
	"context"
	"crypto/md5"
	"fmt"
	"github.com/EdwardShen125/bytedance-simple-server/cmd/user/dal/db"
	"github.com/EdwardShen125/bytedance-simple-server/kitex_gen/user"
	"github.com/EdwardShen125/bytedance-simple-server/pkg/errno"
	"io"
)

type RegisterUserService struct {
	ctx context.Context
}

func NewRegisterUserService(ctx context.Context) *RegisterUserService {
	return &RegisterUserService{
		ctx: ctx,
	}
}

func (s *RegisterUserService) RegisterUser(req *user.UserRegisterRequest) (int64, error) {
	users, err := db.QueryUserByName(s.ctx, req.Username)
	if err != nil {
		return 0, err
	}
	if len(users) != 0 {
		return 0, errno.UserAlreadyExistErr
	}

	h := md5.New()
	if _, err = io.WriteString(h, req.Password); err != nil {
		return 0, err
	}
	password := fmt.Sprintf("%x", h.Sum(nil))

	userId, err := db.UploadUserInfo(s.ctx, req.Username, password)
	if err != nil {
		return 0, err
	}

	return userId, nil

}
