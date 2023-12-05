package data

import (
	"context"

	"gorm.io/gorm"
	"github.com/sjxiang/webook-backend/internal/biz"
)

// func (u *user) Create(ctx context.Context, user *biz.UserDO) error {
// 	err := u.db.WithContext(ctx).Create(u).Error
// 	if me, ok := err.(*mysql.MySQLError); ok {
// 		const uniqueIndexErrNo uint16 = 1062
// UniqueViolation     = "23505"  // 违反唯一约束
// Error 1062 (23000): Duplicate entry
// 		if me.Number == uniqueIndexErrNo {
// 			return errno.ErrUserDuplicate.WithMessage(err.Error())
// 		}

// 		return errno.ErrDatabaseFail.WithMessage(err.Error())
// 	}
// 	return nil
// }

// CreateUser 插入一条 user 记录.
func (ur *userRepo) CreateUser(ctx context.Context, u *biz.User) error {
	user := &UserM{
		Email:    u.Email,
		Password: u.Password,
	}
	
	return ur.storage.WithContext(ctx).Create(user).Error
}


// UserM 对应数据库表结构，可参考，entity、model、PO
type UserM struct {
	gorm.Model
	Username  string    `gorm:"column:username;type:varchar(20)"`
	Password  string    `gorm:"column:password;type:varchar(128);not null"`
	Email     string    `gorm:"column:email;type:varchar(30);idx_email;unique;not null"`
	Mobile    string    `gorm:"column:mobile;type:varchar(11)"`
	Gender    string    `gorm:"column:gender;type:varchar(6);default:male;comment 'female表示女，male表示男'"`
	Role      int       `gorm:"column:role;type:int;default:1;comment '1表示普通用户，2表示管理员'"`
}

// TableName 表名
func (u *UserM) TableName() string {
	return "user"
}

// func (u *user) List(ctx context.Context, orderby []string, opts biz.ListMeta) (*biz.UserDOList, error) {
	
// 	var users *biz.UserDOList
	
// 	// 分页
// 	var limit, offset int
// 	if opts.PageSize == 0 {
// 		limit = 10
// 	} else {
// 		limit = opts.PageSize
// 	}
	
// 	if opts.PageNum > 0 {
// 		offset = (opts.PageNum - 1)*limit
// 	} 

// 	// 排序
// 	query := u.db  // 防止影响全局
// 	for _, val := range orderby {
// 		query = query.Order(val)
// 	}

// 	// 查询
// 	tx := query.Offset(offset).Limit(limit).Find(&users.Items).Count(&users.TotalCount)
// 	if tx.Error != nil {
// 		return nil, errno.ErrDatabaseFail.WithMessage(tx.Error.Error())
// 	}

// 	return users, nil
// }

// func (u *user) GetByMobile(ctx context.Context, mobile string) (*biz.UserDO, error) {
// 	var user biz.UserDO
// 	err := u.db.WithContext(ctx).Where("mobile = ?", mobile).First(&user).Error
// 	if err != nil {
// 		if errors.Is(err, gorm.ErrRecordNotFound) {
// 			return nil, errno.ErrUserNotFound.WithMessage(err.Error())
// 		}
// 		return nil, errno.ErrDatabaseFail.WithMessage(err.Error())
// 	}
// 	return &user, err
// }

// func (u *user) GetById(ctx context.Context, id uint64) (*biz.UserDO, error) {
// 	var user *biz.UserDO
// 	err := u.db.WithContext(ctx).First(user, id).Error
// 	if err != nil && err != gorm.ErrRecordNotFound {
// 		return nil, errno.ErrDatabaseFail.WithMessage(err.Error())
// 	}
// 	if user == nil {
// 		return nil, errno.ErrUserNotFound.WithMessage(err.Error())
// 	} 
// 	return user, err
// }


// func (u *user)Update(ctx context.Context, user *biz.UserDO) error {
// 	tx := u.db.Save(user)
// 	if tx.Error != nil {
// 		return errno.ErrDatabaseFail.WithMessage(tx.Error.Error())
// 	}
// 	return nil
// }