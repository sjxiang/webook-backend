package controller

import (
	"go.uber.org/zap"
	"github.com/sjxiang/webook-backend/internal/biz"
	"github.com/sjxiang/webook-backend/pkg/token"
)


type Controller struct {
	uc         *biz.UserUsecase
	// sc      *biz.SocialUsecase

	tokenMaker token.Maker
	logger     *zap.SugaredLogger
}

func (controller *Controller) ExportTokenMaker() token.Maker {
	return controller.tokenMaker
}

func NewControllerForBackend(uc *biz.UserUsecase, tokenMaker token.Maker, logger *zap.SugaredLogger) *Controller {
	return &Controller{
		uc:         uc,
		logger:     logger,
		tokenMaker: tokenMaker,
	}
}
	