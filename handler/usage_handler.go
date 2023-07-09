package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sfcup/dal"
	"sfcup/response"
	"time"
)

func GetSelfUsage(c *gin.Context) {
	id := c.MustGet("id").(int64)
	result, err := dal.Usage.Where(dal.Usage.UserID.Eq(id)).Find()
	if err != nil {
		response.Send(c, http.StatusInternalServerError, nil, "服务器错误")
		return
	}
	response.Send(c, http.StatusOK, result, "")
}

func GetTotalUsageStatistic(c *gin.Context) {
	// 创建空的切片
	dates := make([]string, 0)
	data := make([]int64, 0)

	// 循环生成最近七天的日期
	for i := 6; i >= 0; i-- {
		offset := -i
		now := time.Now().AddDate(0, 0, offset)
		//时间戳在今天开始到明天开始之间
		dayBegin := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location()).Unix()
		dayEnd := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location()).Add(time.Hour * 24).Unix()
		count, _ := dal.Usage.Where(dal.Usage.CreateAt.Between(dayBegin, dayEnd)).Count()
		formattedDate := now.Format("01-02")
		dates = append(dates, formattedDate)
		data = append(data, count)
	}
	response.Send(c, http.StatusOK, gin.H{"dates": dates, "data": data}, "")
}
