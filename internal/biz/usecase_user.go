package biz

import (
	"context"

	"github.com/sjxiang/webook-backend/internal/xerr"
	"github.com/sjxiang/webook-backend/pkg/util"
)


type User struct {
	ID       int64
	NickName string
	Email    string
	Mobile   string
	Password string
	Intro    string
	Birthday int64 
	Avatar   string
}


func (uc *UserUsecase) Register(ctx context.Context, email, password string) error {
	// 敏感数据
	hashedPassword, err := util.HashPassword(password)
	if err != nil {
		return err
	}

	user := &User{
		Email:    email,
		Password: hashedPassword,
	}
	return uc.ur.CreateUser(ctx, user)
}


func (uc *UserUsecase) Login(ctx context.Context, email, password string) (*User, error) {
	// 查找邮箱
	u , err := uc.ur.GetUserByEmail(ctx, email)
	uc.logger.Info("uc", u, email)
	if err != nil {
		return nil, err
	}

	// 校验密码
	if err := util.CheckPassword(password, u.Password); err != nil {
		return nil, xerr.InvalidUserOrPassword
	}

	return u, nil
}

func (uc *UserUsecase) Profile(ctx context.Context, uid int64) (*User, error) {
	// 查找用户
	u , err := uc.ur.GetUserByID(ctx, uid)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (uc *UserUsecase) Edit(ctx context.Context, uid int64, username, intro string, birthday int64, avatar string) error {
	user := &User{
		ID:       uid,
		NickName: username,
		Birthday: birthday,
		Intro:    intro,
		Avatar:   avatar,
	}
	return uc.ur.UpdateByID(ctx, user)
}


// UserRepo 接口，定义了 data 层需要提供的能力，此接口实现者为 data/repo_user.go 文件中的 userRepo
type UserRepo interface {
	CreateUser(ctx context.Context, user *User) error
	GetUserByEmail(ctx context.Context, email string) (*User, error)
	GetUserByID(ctx context.Context, id int64) (*User, error) 
	UpdateByID(ctx context.Context, user *User) error 
}
