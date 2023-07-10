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
	"time"
)

type GetSelfFilesDTO struct {
	Filename    string `json:"filename"`
	CreateTime  string `json:"create_time"`
	PatientAge  int    `json:"patient_age"`
	PatientName string `json:"patient_name"`
	Status      string `json:"status"`
}

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
	err = dal.File.Create(&model.File{UserID: id, Filename: fileName, PatientName: name, PatientAge: int64(age), Status: "已上传"})
	if err != nil {
		response.Send(c, http.StatusInternalServerError, nil, "服务器错误")
		return
	}
	response.Send(c, http.StatusOK, fileName, "")
}

func GetSelfFiles(c *gin.Context) {
	var files []GetSelfFilesDTO
	id := c.MustGet("id").(int64)
	result, err := dal.File.Where(dal.File.UserID.Eq(id)).Find()
	if err != nil {
		return
	}
	for _, v := range result {
		file := GetSelfFilesDTO{
			Filename:    v.Filename,
			CreateTime:  time.Unix(v.CreateTime, 0).Format("2006-01-02 15:04:05"),
			PatientAge:  int(v.PatientAge),
			PatientName: v.PatientName,
			Status:      v.Status,
		}
		files = append(files, file)
	}
	response.Send(c, http.StatusOK, files, "")
}

func DownloadFile(c *gin.Context) {
	url := c.Param("fileName")
	c.File("file/" + url)
}
