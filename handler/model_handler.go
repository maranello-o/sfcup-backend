package handler

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"sfcup/dal"
	"sfcup/model"
	"sfcup/response"
)

func GetPredictResult(c *gin.Context) {
	id := c.MustGet("id").(int64)
	fileName := c.Param("fileName")
	modelName := c.Param("modelName")
	dal.Usage.Create(&model.Usage{
		UserID: id,
		Model:  modelName,
	})
	//变更：文件提前传到后端了
	//收到文件
	//file, err := c.FormFile("file")
	//if err != nil {
	//	response.Send(c, http.StatusBadRequest, err.Error(), "参数错误")
	//	return
	//}
	// 创建转发请求
	url := "http://10.13.120.32:5000/infer/" + modelName
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		response.Send(c, http.StatusBadRequest, nil, err.Error())
		return
	}
	// 把来自前端的form数据转发到外部API上，将原来的formfile中的文件放入转发请求的body
	formData := &bytes.Buffer{}
	writer := multipart.NewWriter(formData)
	filePart, err := writer.CreateFormFile("file", fileName)
	if err != nil {
		response.Send(c, http.StatusBadRequest, nil, err.Error())
		return
	}
	fileContent, err := os.Open("./file/" + fileName)
	if err != nil {
		response.Send(c, http.StatusBadRequest, nil, err.Error())
		return
	}
	defer fileContent.Close()
	_, err = io.Copy(filePart, fileContent)
	if err != nil {
		response.Send(c, http.StatusBadRequest, nil, err.Error())
		return
	}
	writer.Close()
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Body = io.NopCloser(formData)
	//进行请求
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		response.Send(c, http.StatusBadRequest, nil, err.Error())
		return
	}
	defer resp.Body.Close()
	//预测服务器返回的文件数据
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.String(http.StatusInternalServerError, "Internal error")
		return
	}
	fmt.Println("成功收到外部服务器文件")
	finalFileName := "./file/final-" + fileName
	err = os.WriteFile(finalFileName, body, 0644)
	if err != nil {
		response.Send(c, http.StatusBadRequest, nil, err.Error())
		return
	}
	err = dal.Usage.Create(&model.Usage{Model: modelName, UserID: id})
	if err != nil {
		response.Send(c, http.StatusInternalServerError, nil, "服务器错误")
		return
	}
	//修改文件状态为已分割
	dal.File.Where(dal.File.Filename.Eq(fileName)).Update(dal.File.Status, "已分割")
	response.Send(c, http.StatusOK, finalFileName, "")
	//c.Data(resp.StatusCode, resp.Header.Get("Content-Type"), body)
}
