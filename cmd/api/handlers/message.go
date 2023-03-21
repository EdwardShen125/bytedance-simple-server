package handlers

import (
	"context"
	"strconv"

	"github.com/EdwardShen125/bytedance-simple-server/cmd/api/rpc"
	"github.com/EdwardShen125/bytedance-simple-server/kitex_gen/message"
	"github.com/EdwardShen125/bytedance-simple-server/pkg/errno"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/gin-gonic/gin"
)

func MessageAction(c *gin.Context) {
	token := c.Query("token")
	toUserId := c.Query("to_user_id")
	actionType := c.Query("action_type")
	content := c.Query("content")
	klog.Infof("MessageAction content[%v]\n", content)
	if actionType != "1" {
		SendResponse(c, errno.ParamParseErr)
		return
	}
	toUserId_, err := strconv.ParseInt(toUserId, 10, 64)
	if err != nil {
		SendResponse(c, errno.ParamParseErr)
		return
	}
	at, err := strconv.Atoi(actionType)
	if err != nil {
		SendResponse(c, errno.ParamParseErr)
		return
	}
	actionType_ := int32(at)
	if err != nil {
		SendResponse(c, errno.ParamParseErr)
		return
	}
	req := &message.MessageActionRequest{
		Token:      token,
		ToUserId:   toUserId_,
		ActionType: actionType_,
		Content:    content,
	}
	err = rpc.MessageAction(context.Background(), req)
	if err != nil {
		SendResponse(c, errno.ConvertErr(err))
		return
	}
	SendResponse(c, errno.Success)
}

func MessageChat(c *gin.Context) {
	token := c.Query("token")
	toUserId := c.Query("to_user_id")
	toUserId_, err := strconv.ParseInt(toUserId, 10, 64)
	if err != nil {
		SendChatResponse(c, nil, errno.ParamParseErr)
		return
	}

	req := &message.MessageChatRequest{
		Token:    token,
		ToUserId: toUserId_,
		//PreMsgTime: 0,
	}
	messages, err := rpc.MessageChat(context.Background(), req)
	if err != nil {
		SendChatResponse(c, nil, errno.ConvertErr(err))
		return
	}
	SendChatResponse(c, messages, errno.Success)
}
