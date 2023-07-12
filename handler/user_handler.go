package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sfcup/dal"
	"sfcup/response"
	"time"
)

type getSelfProfileDTO struct {
	ID         int64  `json:"id"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Nick       string `json:"nick"`
	Avatar     string `json:"avatar"`
	CreateTime string `json:"create_time"`
}

func GetSelfProfile(c *gin.Context) {
	id := c.MustGet("id").(int64)
	//必定有一个，因为已登录的用户才能请求
	user, err := dal.User.Where(dal.User.ID.Eq(id)).Take()
	if err != nil {
		response.Send(c, http.StatusBadRequest, err.Error(), "服务器错误，请稍后重试")
		return
	}
	createTime := time.Unix(user.CreateTime, 0).Format("2006-01-02")
	dto := getSelfProfileDTO{ID: id, Email: user.Email, Password: user.Password, Avatar: user.Avatar, Nick: user.Nickname, CreateTime: createTime}
	response.Send(c, http.StatusOK, dto, "")
}
