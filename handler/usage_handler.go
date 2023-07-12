package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sfcup/dal"
	"sfcup/response"
	"time"
)

type GetSelfUsageDTO struct {
	Model string `json:"model"`
	Time  string `json:"time"`
}

func GetSelfUsage(c *gin.Context) {
	var usages []GetSelfUsageDTO
	id := c.MustGet("id").(int64)
	result, err := dal.Usage.Where(dal.Usage.UserID.Eq(id)).Order(dal.Usage.CreateTime.Desc()).Find()
	if err != nil {
		response.Send(c, http.StatusInternalServerError, nil, "服务器错误")
		return
	}
	for _, v := range result {
		usage := GetSelfUsageDTO{Model: v.Model, Time: time.Unix(v.CreateTime, 0).Format("2006-01-02 15:04:05")}
		usages = append(usages, usage)
	}
	response.Send(c, http.StatusOK, usages, "")
}

func GetModelUsageProportion(c *gin.Context) {
	var data []struct {
		Name  string `json:"name"`
		Value int    `json:"value"`
	}
	err := dal.Usage.Select(dal.Usage.ID.Count().As("value"), dal.Usage.Model.As("name")).Group(dal.Usage.Model).Scan(&data)
	if err != nil {
		response.Send(c, http.StatusOK, nil, err.Error())
		return
	}

	response.Send(c, http.StatusOK, data, "")
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
		count, _ := dal.Usage.Where(dal.Usage.CreateTime.Between(dayBegin, dayEnd)).Count()
		formattedDate := now.Format("01-02")
		dates = append(dates, formattedDate)
		data = append(data, count)
	}
	response.Send(c, http.StatusOK, gin.H{"dates": dates, "data": data}, "")
}
