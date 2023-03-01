package pack

import (
	"github.com/EdwardShen125/bytedance-simple-server/cmd/comment/dal/db"
	"github.com/EdwardShen125/bytedance-simple-server/kitex_gen/comment"
	"github.com/EdwardShen125/bytedance-simple-server/pkg/constants"
)

func CommentList(currentId int64, comments []*db.CommentRaw, userMap map[int64]*db.UserRaw, relationMap map[int64]*db.RelationRaw) []*comment.Comment {
	commentList := make([]*comment.Comment, 0)
	for _, commentRaw := range comments {
		commentUser, ok := userMap[commentRaw.UserId]
		if !ok {
			commentUser = &db.UserRaw{
				Name:          "未知用户",
				FollowCount:   0,
				FollowerCount: 0,
			}
			commentUser.ID = 0
		}

		var isFollow = false

		if currentId != -1 {
			_, ok := relationMap[commentRaw.UserId]
			if ok {
				isFollow = true
			}
		}

		commentList = append(commentList, &comment.Comment{
			Id: int64(commentRaw.ID),
			User: &comment.User{
				Id:            int64(commentUser.ID),
				Name:          commentUser.Name,
				FollowCount:   &commentUser.FollowCount,
				FollowerCount: &commentUser.FollowerCount,
				IsFollow:      isFollow,
			},
			Content:    commentRaw.Contents,
			CreateDate: commentRaw.UpdatedAt.Format(constants.TimeFormat),
		})
	}
	return commentList
}

func CommentInfo(commentRaw *db.CommentRaw, user *db.UserRaw) *comment.Comment {
	comment_ := &comment.Comment{
		Id: int64(commentRaw.ID),
		User: &comment.User{
			Id:            int64(user.ID),
			Name:          user.Name,
			FollowCount:   &user.FollowCount,
			FollowerCount: &user.FollowerCount,
			IsFollow:      false,
		},
		Content:    commentRaw.Contents,
		CreateDate: commentRaw.UpdatedAt.Format(constants.TimeFormat),
	}
	return comment_
}