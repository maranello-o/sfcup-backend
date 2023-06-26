package handler

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/okieraised/gonii"
	"io"
	"net/http"
	"sfcup/response"
)

func GetVolume(c *gin.Context) {
	fileHeader, err := c.FormFile("file")
	if err != nil {
		response.Send(c, http.StatusBadRequest, err.Error(), "")
		return
	}
	file, err := fileHeader.Open()
	if err != nil {
		response.Send(c, http.StatusInternalServerError, err.Error(), "")
		return
	}
	fileBytes, err := io.ReadAll(file)
	if err != nil {
		response.Send(c, http.StatusInternalServerError, err.Error(), "")
		return
	}
	rd, err := gonii.NewNiiReader(gonii.WithReadImageReader(bytes.NewReader(fileBytes)), gonii.WithReadRetainHeader(true))
	if err != nil {
		panic(err)
	}
	// Parse the image
	err = rd.Parse()
	if err != nil {
		panic(err)
	}
	data := rd.GetNiiData()
	pixdim := data.PixDim
	//获取体素对应的物理大小 mm^3
	vsize := pixdim[1] * pixdim[2] * pixdim[3]
	volume := map[float64]float64{1.0: 0, 2.0: 0, 3.0: 0, 4.0: 0, 5.0: 0, 6.0: 0, 7.0: 0, 8.0: 0, 9.0: 0, 10.0: 0, 11.0: 0}

	volumeArray, err := data.GetVolume(0)
	if err != nil {
		response.Send(c, http.StatusBadRequest, err.Error(), "")
	}
	for _, x := range volumeArray {
		for _, y := range x {
			for _, z := range y {
				_, ok := volume[z]
				if ok {
					volume[z] += vsize
				}
			}
		}
	}
	c.String(http.StatusOK, fmt.Sprintf("%v", volume))
}
