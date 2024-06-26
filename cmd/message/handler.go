package main

import (
	"context"

	"github.com/EdwardShen125/bytedance-simple-server/cmd/message/pack"
	"github.com/EdwardShen125/bytedance-simple-server/cmd/message/service"
	"github.com/EdwardShen125/bytedance-simple-server/kitex_gen/message"
	"github.com/EdwardShen125/bytedance-simple-server/pkg/errno"
)

type MessageServiceImpl struct{}

func (s *MessageServiceImpl) MessageChat(ctx context.Context, req *message.MessageChatRequest) (r *message.MessageChatResponse, err error) {
	resp := new(message.MessageChatResponse)
	messages, err := service.NewMessageChatService(ctx).MessageChat(req)
	if err != nil {
		resp = pack.BuildMessageChatResp(err)
		return resp, err
	}
	resp = pack.BuildMessageChatResp(errno.Success)
	resp.MessageList = messages
	return resp, nil
}

func (s *MessageServiceImpl) MessageAction(ctx context.Context, req *message.MessageActionRequest) (resp *message.MessageActionResponse, err error) {
	resp = new(message.MessageActionResponse)
	if req.ActionType != 1 {
		resp = pack.BuildMessageActionResp(errno.ParamErr)
		return resp, nil
	}
	_, err = service.NewActionService(ctx).MessageAction(req)
	if err != nil {
		resp = pack.BuildMessageActionResp(err)
		return resp, nil
	}
	resp = pack.BuildMessageActionResp(errno.Success)
	return resp, nil
}
