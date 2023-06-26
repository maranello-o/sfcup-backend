package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"sfcup/dal"
	"sfcup/model"
	"sfcup/response"
	"sfcup/util"
)

var Codes = make(map[string]string, 10)

type registerDTO struct {
	Nick     string `binding:"required" json:"nick"`
	Email    string `binding:"required,email"`
	Password string `binding:"required"`
	Code     string `binding:"required,len=6"`
}

type loginDTO struct {
	Email    string `binding:"required,email" json:"email"`
	Password string `binding:"required" json:"password"`
}

func Login(c *gin.Context) {
	var dto loginDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		response.Send(c, http.StatusBadRequest, err.Error(), "参数错误")
		return
	}

	user, err2 := dal.User.Where(dal.User.Email.Eq(dto.Email)).Take()
	if err2 != nil {
		if errors.Is(err2, gorm.ErrRecordNotFound) {
			response.Send(c, http.StatusBadRequest, nil, "用户不存在")
			return
		}
		response.Send(c, http.StatusBadRequest, err2.Error(), "服务器错误")
		return
	}

	if dto.Password != user.Password {
		response.Send(c, http.StatusBadRequest, nil, "密码错误")
		return
	}
	token, err := util.GenJWT(user.ID)
	if err != nil {
		response.Send(c, http.StatusBadRequest, err.Error(), "Token生成错误")
		return
	}
	response.Send(c, http.StatusOK, token, "")
}

func GenVerificationCode(c *gin.Context) {
	genCode := util.CodeGen()
	email := c.Query("email")

	if err := util.SendCode(email, genCode); err != nil {
		response.Send(c, http.StatusBadRequest, nil, "邮件发送失败")
		return
	}
	// Redis实现限时验证码
	//if err := da.RDB.Set(email, genCode, time.Minute*3).Err(); err != nil {
	//	response.Send(c, http.StatusBadRequest, nil, "验证码存储失败")
	//	return
	//}
	Codes[email] = genCode
	response.Send(c, http.StatusOK, nil, "")
}

func RegisterUser(c *gin.Context) {
	var dto registerDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		response.Send(c, http.StatusBadRequest, nil, "参数错误")
		return
	}
	email := dto.Email
	_, err := dal.User.Where(dal.User.Email.Eq(dto.Email)).Take()
	if err == nil {
		response.Send(c, http.StatusBadRequest, nil, "该邮箱已注册")
		return
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		response.Send(c, http.StatusBadRequest, nil, "服务器错误")
		return
	}

	//trueCode, err := da.RDB.Get(email).Result()
	trueCode := Codes[email]
	if dto.Code != trueCode {
		response.Send(c, http.StatusBadRequest, nil, "验证码错误")
		return
	}

	password := dto.Password

	user := model.User{Email: email, Nickname: dto.Nick, Password: password}
	if err2 := dal.User.Create(&user); err2 != nil {
		response.Send(c, http.StatusBadRequest, nil, "数据库错误")
		return
	}
	delete(Codes, email)
	response.Send(c, http.StatusOK, nil, "")

}
