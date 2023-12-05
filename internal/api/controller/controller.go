package controller

import (
	"go.uber.org/zap"
	"github.com/sjxiang/webook-backend/internal/biz"
)


type Controller struct {
	uc     *biz.UserUsecase
	// sc     *biz.SocialUsecase
	logger *zap.SugaredLogger
}


func NewControllerForBackend(uc *biz.UserUsecase, logger *zap.SugaredLogger) *Controller {
	return &Controller{
		uc:     uc,
		logger: logger,
	}
}
	