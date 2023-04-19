package main

import (
	"context"
	"reflect"
	"testing"
)

func TestMessageServiceImpl_MessageAction(t *testing.T) {
	type args struct {
		ctx context.Context
		req *message.MessageActionRequest
	}
	tests := []struct {
		name     string
		args     args
		wantResp *message.MessageActionResponse
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &MessageServiceImpl{}
			gotResp, err := s.MessageAction(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("MessageAction() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("MessageAction() gotResp = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}

func TestMessageServiceImpl_MessageChat(t *testing.T) {
	type args struct {
		ctx context.Context
		req *message.MessageChatRequest
	}
	tests := []struct {
		name    string
		args    args
		wantR   *message.MessageChatResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &MessageServiceImpl{}
			gotR, err := s.MessageChat(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("MessageChat() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotR, tt.wantR) {
				t.Errorf("MessageChat() gotR = %v, want %v", gotR, tt.wantR)
			}
		})
	}
}
