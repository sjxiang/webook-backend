package biz

import "go.uber.org/zap"


type UserUsecase struct {
// storage
// cache
// drive
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

}


// // UserRepo 接口，定义了 data 层需要提供的能力，此接口实现者为 data/user.go 文件中的 userRepo
// type UserRepo interface {
// 	CreateUser(ctx context.Context, user *User) error
// 	GetUserByEmail(ctx context.Context, email string) (*User, error)
// 	GetUserByUsername(ctx context.Context, username string) (*User, error)
// 	GetUserByID(ctx context.Context, id uint) (*User, error)
// 	UpdateUser(ctx context.Context, user *User) (*User, error)
// }

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