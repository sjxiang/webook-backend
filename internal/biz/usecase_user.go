package biz

import "context"

type User struct {
	Email    string
	Password string 
}


func (uc *UserUsecase) Register(ctx context.Context, email, password string) error {
	user := &User{
		Email:    email,
		Password: password,
	}
	return uc.ur.CreateUser(ctx, user)
}



// UserRepo 接口，定义了 data 层需要提供的能力，此接口实现者为 data/user.go 文件中的 userRepo
type UserRepo interface {
	CreateUser(ctx context.Context, user *User) error
	// GetUserByEmail(ctx context.Context, email string) (*User, error)
	// GetUserByUsername(ctx context.Context, username string) (*User, error)
	// GetUserByID(ctx context.Context, id uint) (*User, error)
	// UpdateUser(ctx context.Context, user *User) (*User, error)
}
