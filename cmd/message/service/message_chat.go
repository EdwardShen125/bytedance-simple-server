package service

import (
	"context"

	"github.com/EdwardShen125/bytedance-simple-server/cmd/message/dal/db"
	"github.com/EdwardShen125/bytedance-simple-server/cmd/message/pack"
	"github.com/EdwardShen125/bytedance-simple-server/kitex_gen/message"
	"github.com/EdwardShen125/bytedance-simple-server/pkg/constants"
	"github.com/EdwardShen125/bytedance-simple-server/pkg/jwt"
)

type MessageChatService struct {
	ctx context.Context
}

func NewMessageChatService(ctx context.Context) *MessageChatService {
	return &MessageChatService{ctx: ctx}
}

func (s *MessageChatService) MessageChat(req *message.MessageChatRequest) ([]*message.Message, error) {
	Jwt := jwt.NewJWT([]byte(constants.SecretKey))
	currentId, _ := Jwt.CheckToken(req.Token)
	messages := make([]*db.MessageRaw, 0)
	messages, err := db.QueryMessageById(s.ctx, currentId, req.ToUserId)
	if err != nil {
		return nil, err
	}
	messages_ := make([]*message.Message, 0)
	messages_ = pack.MessageList(messages)

	return messages_, nil
}
