package controller

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/sjxiang/webook-backend/internal/xerr"
)

// 登录
type LoginReq struct {
	Email    string `json:"email"    validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=32"`
	// VerificationCode string `json:"verification_code" validate:"required,len=6"`
}

func (controller *Controller) Login(ctx *gin.Context) {
	// fetch payload
	var req LoginReq
	if err := json.NewDecoder(ctx.Request.Body).Decode(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "请求参数类型不匹配",
		})
		return
	}

	// validate
	if err := validator.New().Struct(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	// biz handle
	user, err := controller.uc.Login(context.TODO(), req.Email, req.Password)

	if err != nil {
		if errors.Is(err, xerr.InvalidUserOrPassword) {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "用户名或者密码不正确，请重试",
			})
			return
		}

		controller.logger.Errorf("系统异常", "biz", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "系统异常",
		})
		return
	}

	// 设置 session
	s := sessions.Default(ctx)
	s.Clear()
	s.Set("user_id", user.ID)
	s.Set("last_time", time.Now())
	s.Save()

	// feedback
	ctx.JSON(http.StatusOK, gin.H{"message": "登录成功"})
}


// JWT v2
func (controller *Controller) LoginJWT(ctx *gin.Context) {
	// fetch payload
	var req LoginReq
	if err := json.NewDecoder(ctx.Request.Body).Decode(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "请求参数类型不匹配",
		})
		return
	}

	// validate
	if err := validator.New().Struct(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	// biz handle
	user, err := controller.uc.Login(context.TODO(), req.Email, req.Password)

	if err != nil {
		if errors.Is(err, xerr.InvalidUserOrPassword) {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "用户名或者密码不正确，请重试",
			})
			return
		}

		controller.logger.Errorf("系统异常", "biz", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "系统异常",
		})
		return
	}

	// 签发 JWT 
	accessToken, accessPayload, err := controller.tokenMaker.CreateToken(user.ID, user.Email, time.Hour*12)
	if err != nil {
		controller.logger.Errorf("系统异常", "biz", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "系统异常",
		})
		return
	}

	// 可以考虑塞进 header x-jwt-token

	// feedback
	ctx.JSON(http.StatusOK, gin.H{
		"message": "登录成功",
		"data": map[string]string{
			"access_token":            accessToken,
			"access_token_expires_at": accessPayload.ExpiredAt.Local().GoString(),
		},
	})
}





