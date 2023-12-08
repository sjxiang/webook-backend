package controller

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// 注销、退出
func (controller *Controller) Logout(ctx *gin.Context) {
	s := sessions.Default(ctx)
	
	s.Options(sessions.Options{
		MaxAge: -1,
	})
	// s.Clear()
	s.Save()
	
	ctx.JSON(http.StatusOK, gin.H{
		"message": "退出登录成功",
	})
	
}