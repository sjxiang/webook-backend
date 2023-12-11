package controller

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sjxiang/webook-backend/pkg/token"
	"github.com/sjxiang/webook-backend/pkg/util"
)

// 注意，其它字段，尤其是密码、邮箱和手机
// 修改都要通过别的手段
// 邮箱和手机都要验证
// 密码更加不用多说了
type editReq struct {
	NickName string `json:"nickname" binding:"required,gte=8,lte=30"`
	Birthday string `json:"birthday" binding:"required,len=10"`  // 1997-09-12
	Intro    string `json:"intro" binding:"required,min=1,max=1000"`
	Avatar   string `json:"avatar" binding:"required"`
}

// 编辑个人基本信息
func (controller *Controller) Edit(ctx *gin.Context) {
	var req editReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": util.Translate(err),
		})
		return
	}

	birthday, err := time.Parse(time.DateOnly, req.Birthday)
	if err != nil {
		// 也就是说，我们其实并没有直接校验具体的格式
		// 而是如果你能转化过来，那就说明没问题
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "日期格式不对",
		})
		return
	}

	// fetch some params
	uid := ctx.MustGet("user_id").(int64)

	err = controller.uc.Edit(context.Background(), uid, req.NickName, req.Intro, birthday.Unix(), req.Avatar)
	if err != nil {
		controller.logger.Errorf("系统异常", "biz", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "系统异常",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "补充资料成功",
	})
}

func (controller *Controller) EditJWT(ctx *gin.Context) {
	var req editReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": util.Translate(err),
		})
		return
	}

	birthday, err := time.Parse(time.DateOnly, req.Birthday)
	if err != nil {
		// 也就是说，我们其实并没有直接校验具体的格式
		// 而是如果你能转化过来，那就说明没问题
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "日期格式不对",
		})
		return
	}

	// fetch some params
	
	// claims, exists := ctx.Get("authorization_payload")
	// if !exists {
	// 	ctx.JSON(http.StatusInternalServerError, gin.H{
	// 		"message": "系统异常",
	// 	})
	// 	return
	// }
	// payload, ok := claims.(*token.Payload)
	// if !ok {
	// 	ctx.JSON(http.StatusInternalServerError, gin.H{
	// 		"message": "系统异常",
	// 	})
	// 	return
	// }

	payload := ctx.MustGet("authorization_payload").(*token.Payload)

	// biz handle
	err = controller.uc.Edit(context.Background(), payload.ID, req.NickName, req.Intro, birthday.Unix(), req.Avatar)
	if err != nil {
		controller.logger.Errorf("系统异常", "biz", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "系统异常",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "补充资料成功",
	})
}