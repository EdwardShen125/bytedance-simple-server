package handlers

import (
	"context"
	"strconv"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/EdwardShen125/bytedance-simple-server/cmd/api/rpc"
	"github.com/EdwardShen125/bytedance-simple-server/kitex_gen/user"
	"github.com/EdwardShen125/bytedance-simple-server/pkg/errno"
	"github.com/gin-gonic/gin"
)

func UserRegister(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	if length := len(username); length <= 0 || length > 32 {
		SendResponse(c, errno.UserNameValidationErr)
		return
	}

	if length := len(password); length <= 0 || length > 32 {
		SendResponse(c, errno.PasswordValidationErr)
		return
	}

	userId, token, err := rpc.UserRegister(context.Background(), &user.UserRegisterRequest{
		Username: username,
		Password: password,
	})
	if err != nil {
		SendResponse(c, errno.ConvertErr(err))
		return
	}
	SendUserResponse(c, errno.Success, userId, token)
}

func UserInfo(c *gin.Context) {
	userIdStr := c.Query("user_id")
	token := c.DefaultQuery("token", "")

	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		SendResponse(c, errno.ParamParseErr)
		return
	}

	klog.Infof("UserInfo userId[%v] token[%v]\n", userId, token)

	user_, err := rpc.UserInfo(context.Background(), &user.UserRequest{
		UserId: userId,
		Token:  token,
	})
	if err != nil {
		SendResponse(c, errno.ConvertErr(err))
		return
	}

	klog.Infof("return user_ %+v\n", user_)
	SendUserInfoResponse(c, errno.Success, user_)
}
