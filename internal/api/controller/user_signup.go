package controller

import (
	"context"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sjxiang/webook-backend/internal/xerr"
	"github.com/sjxiang/webook-backend/pkg/util"
)

// 注册
type signupReq struct {
	Email           string `json:"email" binding:"required,email"`
	Password        string `json:"password" binding:"required,gte=8,lte=30"`
	ConfirmPassword string `json:"confirm_password" binding:"required,gte=8,lte=30"`
}


func (controller *Controller) Signup(ctx *gin.Context) {
	// fetch payload
	var req signupReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": util.Translate(err),
		})
		return
	}

	// validate
	if req.Password != req.ConfirmPassword {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "两次输入密码不对",
		})
		return
	}
	minSize, digit, special, letter := util.ValidatePasswordV1(req.Password)
	if !minSize || !digit || !special || !letter {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "密码必须包含数字、字母、特殊字符，并且长度不能小于 8 位",
		})
		return
	}

	// biz handle
	err := controller.uc.Register(context.Background(), req.Email, req.Password)
	if err != nil {
		if errors.Is(err, xerr.UserDuplicateEmail) {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "邮箱冲突",
			})
			return
		}

		controller.logger.Errorf("系统异常", "biz", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "系统异常",
		})
		return
	}


	// feedback
	ctx.JSON(http.StatusOK, gin.H{"message": "注册成功"})
}
