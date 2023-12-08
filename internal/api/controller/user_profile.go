package controller

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sjxiang/webook-backend/pkg/token"
)

// 用户资料详情
func (controller *Controller) Profile(ctx *gin.Context) {
	uid := ctx.MustGet("user_id").(int64)

	controller.logger.Info("用户详情资料", "biz", uid)
	
	user, err := controller.uc.Profile(context.Background(), uid)
	if err != nil {
		// 按照道理来说，这边 id 对应的数据肯定存在，所以要是没找到，
		// 那就说明是系统出了问题。
		controller.logger.Errorf("系统异常", "biz", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "系统异常",
		})
		return
	}
	
	ctx.JSON(http.StatusOK, user)
}


func (controller *Controller) ProfileJWT(ctx *gin.Context) {
	// fetch some params
	payload := ctx.MustGet("authorization_payload").(*token.Payload)
	
	user, err := controller.uc.Profile(context.Background(), payload.ID)
	if err != nil {
		// 按照道理来说，这边 id 对应的数据肯定存在，所以要是没找到，
		// 那就说明是系统出了问题。
		controller.logger.Errorf("系统异常", "biz", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "系统异常",
		})
		return
	}
	
	ctx.JSON(http.StatusOK, user)
}

