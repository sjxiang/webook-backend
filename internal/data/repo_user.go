package data

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/sjxiang/webook-backend/internal/biz"
	"github.com/sjxiang/webook-backend/internal/xerr"
	"gorm.io/gorm"
)

// Create 插入一条 user 记录.
func (ur *userRepo) CreateUser(ctx context.Context, u *biz.User) error {
	user := &UserM{
		Email:    u.Email,
		Password: u.Password,
	}

	err := ur.storage.WithContext(ctx).Create(user).Error
	if err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			// 违反唯一索引约束 Error 1062 (23000): Duplicate entry
			const uniqueViolation uint16 = 1062  
			if mysqlErr.Number == uniqueViolation {
				return xerr.UserDuplicateEmail
			}
		}
		
		// 其它
		return err
	}
	
	return nil
}

func (ur *userRepo) GetUserByEmail(ctx context.Context, email string) (*biz.User, error) {
	
	// var u UserM
	// panic _ invalid value, should be pointer to struct or slice
	// 期望将查询结果存储在一个有效的结构体指针中，nil

	u := new(UserM)
	err := ur.storage.WithContext(ctx).Where("email = ?", email).First(u).Error

	// 检查 ErrRecordNotFound 错误
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, xerr.InvalidUserOrPassword
	}
	if err != nil {
		return nil, err
	}

	return &biz.User{
		Email:    u.Email,
		Password: u.Password,
	}, nil
}

func (ur *userRepo) GetUserByID(ctx context.Context, id int64) (*biz.User, error) {
	u := new(UserM)
	err := ur.storage.WithContext(ctx).Where("id = ?", id).First(u).Error

	// 检查 ErrRecordNotFound 错误
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, xerr.UserNotFound
	}
	if err != nil {
		return nil, err
	}

	return &biz.User{
		NickName: *u.Username,
		Email:    u.Email,
		// Password: u.Password,
		Intro:    u.Intro,
		Birthday: u.Birthday,
	}, nil
}


func (ur *userRepo)UpdateByID(ctx context.Context, user *biz.User) error {
	u := new(UserM)
	// 这种写法依赖于 GORM 的零值和主键更新特性
	// Update 非零值 WHERE id = ?
	return ur.storage.WithContext(ctx).Model(u).Where("id = ?", user.ID).
		Updates(map[string]any{
			"username": user.NickName,
			"nickname": user.Birthday,
			"intro":    user.Intro,
		}).Error
}

// 一般都是显式指定更新条件、更新字段也尽可能指定，绝对不依赖于默认行为
// 默认行为对后面的维护者很不好
// 尤其是依赖于更新非零值的特性，你看代码是不知道哪些字段是零值，哪些字段是非零值



// UserM 对应数据库表结构，可参考，entity、model、PO
type UserM struct {
	ID        int64             `gorm:"primaryKey,autoIncrement"`
	// 代表这是一个可以为 NULL 的列
	Username  *string           `gorm:"type=varchar(128)"`
	Password  string            `gorm:"type=varchar(128);not null"`
	Email     string            `gorm:"unique;not null"`
	Mobile    string            `gorm:"type=varchar(11)"`
	Intro     string            `gorm:"type=varchar(4096)"`
	// YYYY-MM-DD，UTC 0 的毫秒数
	Birthday  int64             `gorm:"type=varchar(4096)"`
	// 代表这是一个可以为 NULL 的列
	Avatar    sql.NullString 
	Gender    string            `gorm:"column:gender;type:varchar(6);default:male;comment 'female表示女，male表示男'"`
	Role      int               `gorm:"column:role;type:int;default:1;comment '1表示普通用户，2表示管理员'"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt    `gorm:"index"`
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



