package biz


// // 领域模型
// type UserDO struct {
// 	gorm.Model
// 	Mobile   string     `gorm:"idx_mobile;unique;type:varchar(11);not null"`
// 	Password string     `gorm:"type:varchar(64);not null"`
// 	NickName string     `gorm:"type:varchar(20)"`
// 	Birthday *time.Time `gorm:"type:datetime"`
// 	Gender   string     `gorm:"column:gender;type:varchar(6);default:male;comment 'female表示女，male表示男'"`
// 	Role     int        `gorm:"column:role;type:int;default:1;comment '1表示普通用户，2表示管理员'"`
// }

// func (u *UserDO) TableName() string {
// 	return "user"
// }


// type UserDOList struct {
// 	TotalCount int64      `json:"total_count,omitempty"`  // 总数
// 	Items      []*UserDO  `json:"data"`                   // 数据
// }

// type ListMeta struct {
// 	PageNum  int
// 	PageSize int
// }

// type UserRepo interface {
// 	// 用户列表
// 	List(ctx context.Context,  orderby []string, opts ListMeta) (*UserDOList, error)

// 	// 通过手机号查询用户
// 	GetByMobile(ctx context.Context, mobile string) (*UserDO, error)
	
// 	// 通过用户 id 查询用户
// 	GetById(ctx context.Context, id uint64) (*UserDO, error)

// 	// 创建用户
// 	Create(ctx context.Context, user *UserDO) error

// 	// 更新用户
// 	Update(ctx context.Context, user *UserDO) error
// }

/*

命名很烦恼
	Insert、Create、Add
	Find、Get

有数据访问，一定要有 error

参数最好有 ctx

涉及外键，很难评价

*/


