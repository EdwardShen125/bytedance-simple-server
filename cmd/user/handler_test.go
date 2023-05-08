package main

import (
	"context"
	"reflect"
	"testing"

	"github.com/EdwardShen125/bytedance-simple-server/kitex_gen/user"
)

func TestUserServiceImpl_UserInfo(t *testing.T) {
	type args struct {
		ctx context.Context
		req *user.UserRequest
	}
	tests := []struct {
		name     string
		args     args
		wantResp *user.UserResponse
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &UserServiceImpl{}
			gotResp, err := s.UserInfo(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("UserInfo() gotResp = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}

func TestUserServiceImpl_UserLogin(t *testing.T) {
	type args struct {
		ctx context.Context
		req *user.UserLoginRequest
	}
	tests := []struct {
		name     string
		args     args
		wantResp *user.UserLoginResponse
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &UserServiceImpl{}
			gotResp, err := s.UserLogin(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserLogin() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("UserLogin() gotResp = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}

func TestUserServiceImpl_UserRegister(t *testing.T) {
	type args struct {
		ctx context.Context
		req *user.UserRegisterRequest
	}
	tests := []struct {
		name     string
		args     args
		wantResp *user.UserRegisterResponse
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &UserServiceImpl{}
			gotResp, err := s.UserRegister(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserRegister() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("UserRegister() gotResp = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}
