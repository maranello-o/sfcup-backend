package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/okieraised/gonii"
	"net/http"
)

func GetVolume(c *gin.Context) {
	filePath := "file/test2.nii.gz"
	rd, err := gonii.NewNiiReader(gonii.WithReadImageFile(filePath), gonii.WithReadRetainHeader(true))
	if err != nil {
		panic(err)
	}
	// Parse the image
	err = rd.Parse()
	if err != nil {
		panic(err)
	}
	pixdim := rd.GetNiiData().PixDim
	//获取体素
	vsize := pixdim[1] * pixdim[2] * pixdim[3]
	c.String(http.StatusOK, fmt.Sprintf("%f", vsize))
}
