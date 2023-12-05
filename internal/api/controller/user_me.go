package controller

import "github.com/gin-gonic/gin"

// 用户详情
type ProfileRequest struct {
	Email string `json:"email"`
}

// Me
func (controller *Controller) Me(ctx *gin.Context) {

	
}

// func (c *UserHandler) Profile(ctx *gin.Context) {

// 	sess := sessions.Default(ctx)
// 	id := sess.Get(userIdKey).(int64)
// 	u, err := c.svc.Profile(ctx, id)
// 	if err != nil {
// 		// 按照道理来说，这边 id 对应的数据肯定存在，所以要是没找到，
// 		// 那就说明是系统出了问题。
// 		ctx.String(http.StatusOK, "系统错误")
// 		return
// 	}
// 	ctx.JSON(http.StatusOK, Profile{
// 		Email: u.Email,
// 	})
// }