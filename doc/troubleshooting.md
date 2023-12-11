

# 排错集锦

## 1
位置
    (*userRepo).GetUserByID: NickName: *u.Username,
错误
    runtime error: invalid memory address or nil pointer dereference
根因
    Username 是 *string 类型，如果为空，传 nil 就崩
补救措施
    UserM Username字段换成 sql.NullString
    

## 2
位置
    logger.Infow("用户详情", uid)
错误
    zap error Ignored key without a value. 
根因
    日志字段是键值对出现
补救措施
    logger.Info("something happened", zap.Any("key")) 或者 logger.Infow("用户详情", "biz", uid)


## 3
位置
    (*userRepo).GetUserByEmail: var u UserM
错误
    runtime error: invalid memory address or nil pointer dereference
根因
    u 是个空指针，传进 First(&u) 查询，期望将查询结果存储在一个有效的结构体指针中
补救措施
    u := new(UserM)
    前者，仅仅声明变量类型，零值 nil；后者，分配了内存空间，是个有效的结构体指针


## 4


## 5


## 6


## 7

