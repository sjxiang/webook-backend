package token

import (
	"errors"
	"time"
)

var (
	ErrInvalidToken = errors.New("token is invalid")   // 无效，不合规
	ErrExpiredToken = errors.New("token has expired")  // 过期
)

// 载体
type Payload struct {
	ID        int64     `json:"id"`
	Email     string    `json:"email"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`  // 时间戳
}

func NewPayload(id int64, email string, duration time.Duration) *Payload {
	payload := &Payload{
		ID:        id,
		Email:     email,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}
	return payload
}

// 校验 payload 字段是否有效（细分下是不是过期了？还是其它问题导致的无效）
// 等到 jwt.ParseWithClaims 调用时，内部使用
func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpiredAt) {
		return ErrExpiredToken
	}
	return nil
}


/*


type Claims interface {
	Valid() error
}

实现 jwt.Claims 接口

 */
