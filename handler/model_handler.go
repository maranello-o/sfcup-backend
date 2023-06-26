package handler

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"io"
	"mime/multipart"
	"net/http"
	"sfcup/response"
)

func GetPredictResult(c *gin.Context) {
	modelName := c.Param("modelName")
	//收到文件
	file, err := c.FormFile("file")
	if err != nil {
		response.Send(c, http.StatusBadRequest, err.Error(), "参数错误")
		return
	}
	// 创建转发请求
	url := "http://10.13.120.37:5000/infer/" + modelName
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		response.Send(c, http.StatusBadRequest, err.Error(), "")
		return
	}
	// 把来自前端的form数据转发到外部API上，将原来的formfile中的文件放入转发请求的body
	formData := &bytes.Buffer{}
	writer := multipart.NewWriter(formData)
	filePart, err := writer.CreateFormFile("file", file.Filename)
	if err != nil {
		response.Send(c, http.StatusBadRequest, err.Error(), "")
		return
	}
	fileContent, err := file.Open()
	if err != nil {
		response.Send(c, http.StatusBadRequest, err.Error(), "")
		return
	}
	defer fileContent.Close()
	_, err = io.Copy(filePart, fileContent)
	if err != nil {
		response.Send(c, http.StatusBadRequest, err.Error(), "")
		return
	}
	writer.Close()
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Body = io.NopCloser(formData)
	//进行请求
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		response.Send(c, http.StatusBadRequest, err.Error(), "")
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.String(http.StatusInternalServerError, "Internal error")
		return
	}
	c.Data(resp.StatusCode, resp.Header.Get("Content-Type"), body)
}
