package service

import (
	"context"
	"errors"
	"time"

	"github.com/EdwardShen125/bytedance-simple-server/cmd/message/dal/db"
	"github.com/EdwardShen125/bytedance-simple-server/cmd/message/pack"
	"github.com/EdwardShen125/bytedance-simple-server/kitex_gen/message"
	"github.com/EdwardShen125/bytedance-simple-server/pkg/constants"
	"github.com/EdwardShen125/bytedance-simple-server/pkg/jwt"
)

type MessageActionService struct {
	ctx context.Context
}

func NewActionService(ctx context.Context) *MessageActionService {
	return &MessageActionService{ctx: ctx}
}

func (s *MessageActionService) MessageAction(req *message.MessageActionRequest) (*message.Message, error) {
	Jwt := jwt.NewJWT([]byte(constants.SecretKey))
	currentId, err := Jwt.CheckToken(req.Token)
	if err != nil {
		return nil, err
	}
	if req.ActionType == 1 {
		req, err := SendMessage(currentId, req, s.ctx)
		if err != nil {
			return nil, err
		}
		return req, nil
	}
	return nil, errors.New("ActionType err")
}

func SendMessage(userId int64, req *message.MessageActionRequest, ctx context.Context) (*message.Message, error) {
	messageRaw := &db.MessageRaw{
		FromUserID: userId,
		ToUserID:   req.ToUserId,
		Content:    req.Content,
		CreateTime: time.Now(),
	}
	err := db.PostMessage(ctx, *messageRaw)
	if err != nil {
		return nil, err
	}
	message_ := pack.Chat(messageRaw)
	return &message_, nil
}
