package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path"
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

func ChangeAvatar(c *gin.Context) {
	id := c.MustGet("id").(int64)
	file, err := c.FormFile("file")
	if err != nil {
		response.Send(c, http.StatusBadRequest, err.Error(), "传输的文件错误")
		return
	}
	fileName := file.Filename
	dst := path.Join("./file", fileName)
	err2 := os.MkdirAll("file", 0666)
	if err2 != nil {
		fmt.Println(err2)
		response.Send(c, http.StatusBadRequest, nil, "文件保存错误")
		return
	}
	err = c.SaveUploadedFile(file, dst)
	if err != nil {
		response.Send(c, http.StatusBadRequest, nil, "文件保存错误")
		return
	}
	dal.User.Where(dal.User.ID.Eq(id)).Update(dal.User.Avatar, "https://sfcup-backend-production.up.railway.app/file/"+fileName)
	response.Send(c, http.StatusOK, nil, "")
}
