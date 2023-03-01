package handlers

import (
	"context"
	"github.com/EdwardShen125/bytedance-simple-server/cmd/api/rpc"
	"github.com/EdwardShen125/bytedance-simple-server/kitex_gen/feed"
	"github.com/EdwardShen125/bytedance-simple-server/pkg/errno"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

func Feed(c *gin.Context) {

	token := c.DefaultQuery("token", "")
	defaultTime := time.Now().UnixMilli()
	defaultTimeStr := strconv.Itoa(int(defaultTime))
	latestTimeStr := c.DefaultQuery("latest_time", defaultTimeStr)

	//处理传入参数
	latestTime, err := strconv.ParseInt(latestTimeStr, 10, 64)
	if err != nil {
		SendResponse(c, err)
		return
	}

	req := &feed.FeedRequest{LatestTime: &latestTime, Token: &token}
	video, nextTime, err := rpc.Feed(context.Background(), req)
	if err != nil {
		SendResponse(c, err)
		return
	}
	SendFeedResponse(c, errno.Success, video, nextTime)
}
