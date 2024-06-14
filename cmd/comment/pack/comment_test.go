package pack

import (
	"github.com/EdwardShen125/bytedance-simple-server/cmd/comment/dal/db"
	"github.com/EdwardShen125/bytedance-simple-server/kitex_gen/comment"
	"reflect"
	"testing"
)

func TestCommentList(t *testing.T) {
	type args struct {
		currentId   int64
		comments    []*db.CommentRaw
		userMap     map[int64]*db.UserRaw
		relationMap map[int64]*db.RelationRaw
	}
	tests := []struct {
		name string
		args args
		want []*comment.Comment
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CommentList(tt.args.currentId, tt.args.comments, tt.args.userMap, tt.args.relationMap); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CommentList() = %v, want %v", got, tt.want)
			}
		})
	}
}
