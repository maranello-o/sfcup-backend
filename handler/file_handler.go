package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path"
	"sfcup/response"
)

func UploadFile(c *gin.Context) {
	file, err := c.FormFile("file")
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
	response.Send(c, http.StatusOK, fileName, "")
}

//func DownloadFile(c *gin.Context) {
//	url := c.Param("fileUrl")
//	c.File("file/" + url)
//}
