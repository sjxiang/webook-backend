package biz

import "go.uber.org/zap"


type UserUsecase struct {
	// token
	ur     UserRepo
	logger *zap.SugaredLogger
}

func NewUserUsecase(ur UserRepo, logger *zap.SugaredLogger) *UserUsecase {
	return &UserUsecase{
		ur:     ur,
		logger: logger,
	}
}

type SocialUsecase struct {
	// 
	logger *zap.SugaredLogger
}



// type ProfileRepo interface {
// 	GetProfile(ctx context.Context, username string) (*Profile, error)
// 	FollowUser(ctx context.Context, currentUserID uint, followingID uint) error
// 	UnfollowUser(ctx context.Context, currentUserID uint, followingID uint) error
// 	GetUserFollowingStatus(ctx context.Context, currentUserID uint, userIDs []uint) (following []bool, err error)
// }

// // UserUsecase 用户领域结构体，可以包含多个与用户业务相关的 repo
// type UserUsecase struct {
// 	ur   UserRepo
// 	pr   ProfileRepo
// 	jwtc *conf.JWT

// 	log *log.Helper
// }

// type Profile struct {
// 	ID        uint
// 	Username  string
// 	Bio       string
// 	Image     string
// 	Following bool
// }

// // NewUserUsecase 用户领域构造方法
// func NewUserUsecase(ur UserRepo,
// 	pr ProfileRepo, logger log.Logger, jwtc *conf.JWT) *UserUsecase {
// 	return &UserUsecase{ur: ur, pr: pr, jwtc: jwtc, log: log.NewHelper(logger)}
// }