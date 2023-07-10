package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path"
	"sfcup/dal"
	"sfcup/model"
	"sfcup/response"
	"strconv"
)

func UploadFile(c *gin.Context) {
	file, err := c.FormFile("file")
	name := c.Query("name")
	ageStr := c.Query("age")
	if name == "" || ageStr == "" {
		response.Send(c, http.StatusBadRequest, nil, "参数错误")
		return
	}
	age, err := strconv.Atoi(ageStr)
	if err != nil {
		response.Send(c, http.StatusBadRequest, nil, "参数错误")
		return
	}
	id := c.MustGet("id").(int64)
	if err != nil {
		response.Send(c, http.StatusBadRequest, nil, "参数错误")
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
	err3 := c.SaveUploadedFile(file, dst)
	if err3 != nil {
		fmt.Println(err3)
		response.Send(c, http.StatusBadRequest, nil, "文件保存错误")
		return
	}
	//创建数据库记录
	err = dal.File.Create(&model.File{UserID: id, Filename: fileName, PatientName: name, PatientAge: int64(age)})
	if err != nil {
		response.Send(c, http.StatusInternalServerError, nil, "服务器错误")
		return
	}
	response.Send(c, http.StatusOK, fileName, "")
}

func DownloadFile(c *gin.Context) {
	url := c.Param("fileName")
	c.File("file/" + url)
}
