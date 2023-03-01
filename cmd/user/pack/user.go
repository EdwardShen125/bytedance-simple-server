package pack

import (
	"github.com/EdwardShen125/bytedance-simple-server/cmd/user/dal/db"
	"github.com/EdwardShen125/bytedance-simple-server/kitex_gen/user"
)

func UserInfo(userRaw *db.UserRaw, isFollow bool) *user.User {
	userInfo := &user.User{
		Id:            int64(userRaw.ID),
		Name:          userRaw.Name,
		FollowCount:   &userRaw.FollowCount,
		FollowerCount: &userRaw.FollowerCount,
		IsFollow:      isFollow,
	}
	return userInfo
}
