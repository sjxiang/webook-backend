package controller

import (
	"context"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sjxiang/webook-backend/internal/xerr"
)

// 用户资料详情
func (controller *Controller) Profile(ctx *gin.Context) {
	uid := ctx.MustGet("user_id").(int64)

	user, err := controller.uc.Profile(context.TODO(), uid)
	if err != nil {
		if errors.Is(err, xerr.UserNotFound) {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "无用户记录",  // 不大可能
			})
			return
		}

		controller.logger.Errorf("系统异常", "biz", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "系统异常",
		})
		return
	}
	
	ctx.JSON(http.StatusOK, user)
}


// // ProfileJWT 用户详情, JWT 版本
// func (c *UserHandler) ProfileJWT(ctx *gin.Context) {
// 	type Profile struct {
// 		Email    string
// 		Phone    string
// 		Nickname string
// 		Birthday string
// 		AboutMe  string
// 	}
// 	uc := ctx.MustGet("user").(ijwt.UserClaims)
// 	u, err := c.svc.Profile(ctx, uc.Id)
// 	if err != nil {
// 		// 按照道理来说，这边 id 对应的数据肯定存在，所以要是没找到，
// 		// 那就说明是系统出了问题。
// 		ctx.String(http.StatusOK, "系统错误")
// 		return
// 	}
// 	ctx.JSON(http.StatusOK, Profile{
// 		Email:    u.Email,
// 		Phone:    u.Phone,
// 		Nickname: u.Nickname,
// 		Birthday: u.Birthday.Format(time.DateOnly),
// 		AboutMe:  u.AboutMe,
// 	})
// }
