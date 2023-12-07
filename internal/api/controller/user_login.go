package controller

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/sjxiang/webook-backend/internal/xerr"
)

// 登录
type LoginReq struct {
	Email    string `json:"email"    validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=32"`
	// VerificationCode string `json:"verification_code" validate:"required,len=6"`
}

func (controller *Controller) Login(ctx *gin.Context) {
	// fetch payload
	var req LoginReq
	if err := json.NewDecoder(ctx.Request.Body).Decode(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "请求参数类型不匹配",
		})
		return
	}

	// validate
	if err := validator.New().Struct(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	// biz handle
	user, err := controller.uc.Login(context.TODO(), req.Email, req.Password)

	if err != nil {
		if errors.Is(err, xerr.InvalidUserOrPassword) {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "用户名或者密码不正确，请重试",
			})
			return
		}

		controller.logger.Errorf("系统异常", "biz", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "系统异常",
		})
		return
	}

	// 日志成对 Ignored key without a value. 
	controller.logger.Infow("用户登录", "biz", user)

	// 设置 session（把碗掏出来，要饭）
	s := sessions.Default(ctx)
	s.Clear()
	s.Set("user_id", user.ID)
	s.Save()

	// feedback
	ctx.JSON(http.StatusOK, gin.H{"message": "登录成功"})
}



// func (c *UserHandler) RegisterRoutes(server *gin.Engine) {
// 	// 直接注册
// 	//server.POST("/users/signup", c.SignUp)
// 	//server.POST("/users/login", c.Login)
// 	//server.POST("/users/edit", c.Edit)
// 	//server.GET("/users/profile", c.Profile)

// 	// 分组注册
// 	ug := server.Group("/users")
// 	ug.POST("/signup", ginx.WrapReq[SignUpReq](c.SignUp))
// 	// session 机制
// 	//ug.POST("/login", c.Login)
// 	// JWT 机制
// 	ug.POST("/login", c.LoginJWT)
// 	ug.POST("/logout", c.Logout)
// 	ug.POST("/edit", c.Edit)
// 	//ug.GET("/profile", c.Profile)
// 	ug.GET("/profile", c.ProfileJWT)
// 	ug.POST("/login_sms/code/send", c.SendSMSLoginCode)
// 	ug.POST("/login_sms", c.LoginSMS)
// 	ug.POST("/refresh_token", c.RefreshToken)
// }

// func (c *UserHandler) RefreshToken(ctx *gin.Context) {
// 	// 假定长 token 也放在这里
// 	tokenStr := c.ExtractTokenString(ctx)
// 	var rc ijwt.RefreshClaims
// 	token, err := jwt.ParseWithClaims(tokenStr, &rc, func(token *jwt.Token) (interface{}, error) {
// 		return ijwt.RefreshTokenKey, nil
// 	})
// 	// 这边要保持和登录校验一直的逻辑，即返回 401 响应
// 	if err != nil || token == nil || !token.Valid {
// 		ctx.JSON(http.StatusUnauthorized, Result{Code: 4, Msg: "请登录"})
// 		return
// 	}

// 	// 校验 ssid
// 	err = c.CheckSession(ctx, rc.Ssid)
// 	if err != nil {
// 		// 系统错误或者用户已经主动退出登录了
// 		// 这里也可以考虑说，如果在 Redis 已经崩溃的时候，
// 		// 就不要去校验是不是已经主动退出登录了。
// 		ctx.AbortWithStatus(http.StatusUnauthorized)
// 		return
// 	}

// 	err = c.SetJWTToken(ctx, rc.Ssid, rc.Id)
// 	if err != nil {
// 		ctx.JSON(http.StatusUnauthorized, Result{Code: 4, Msg: "请登录"})
// 		return
// 	}
// 	ctx.JSON(http.StatusOK, Result{Msg: "刷新成功"})
// }

// func (c *UserHandler) LoginSMS(ctx *gin.Context) {
// 	type Req struct {
// 		Phone string `json:"phone"`
// 		Code  string `json:"code"`
// 	}
// 	var req Req
// 	if err := ctx.Bind(&req); err != nil {
// 		return
// 	}
// 	ok, err := c.codeSvc.Verify(ctx, bizLogin, req.Phone, req.Code)
// 	if err != nil {
// 		ctx.JSON(http.StatusOK, Result{Code: 5, Msg: "系统异常"})
// 		zap.L().Error("用户手机号码登录失败", zap.Error(err))
// 		return
// 	}
// 	if !ok {
// 		ctx.JSON(http.StatusOK, Result{Code: 4, Msg: "验证码错误"})
// 		return
// 	}

// 	// 验证码是对的
// 	// 登录或者注册用户
// 	u, err := c.svc.FindOrCreate(ctx, req.Phone)
// 	if err != nil {
// 		ctx.JSON(http.StatusOK, Result{Code: 4, Msg: "系统错误"})
// 		return
// 	}
// 	// 用 uuid 来标识这一次会话
// 	ssid := uuid.New().String()
// 	err = c.SetJWTToken(ctx, ssid, u.Id)
// 	if err != nil {
// 		ctx.JSON(http.StatusOK, Result{Msg: "系统错误"})
// 		return
// 	}
// 	ctx.JSON(http.StatusOK, Result{Msg: "登录成功"})
// }

// // SendSMSLoginCode 发送短信验证码
// func (c *UserHandler) SendSMSLoginCode(ctx *gin.Context) {
// 	type Req struct {
// 		Phone string `json:"phone"`
// 	}
// 	var req Req
// 	if err := ctx.Bind(&req); err != nil {
// 		return
// 	}
// 	// 你也可以用正则表达式校验是不是合法的手机号
// 	if req.Phone == "" {
// 		ctx.JSON(http.StatusOK, Result{Code: 4, Msg: "请输入手机号码"})
// 		return
// 	}
// 	err := c.codeSvc.Send(ctx, bizLogin, req.Phone)
// 	switch err {
// 	case nil:
// 		ctx.JSON(http.StatusOK, Result{Msg: "发送成功"})
// 	case service.ErrCodeSendTooMany:
// 		ctx.JSON(http.StatusOK, Result{Code: 4, Msg: "短信发送太频繁，请稍后再试"})
// 	default:
// 		ctx.JSON(http.StatusOK, Result{Code: 5, Msg: "系统错误"})
// 		// 要打印日志
// 		return
// 	}
// }


// // LoginJWT 用户登录接口，使用的是 JWT，如果你想要测试 JWT，就启用这个
// func (c *UserHandler) LoginJWT(ctx *gin.Context) {
// 	type LoginReq struct {
// 		Email    string `json:"email"`
// 		Password string `json:"password"`
// 	}

// 	var req LoginReq
// 	// 当我们调用 Bind 方法的时候，如果有问题，Bind 方法已经直接写响应回去了
// 	if err := ctx.Bind(&req); err != nil {
// 		return
// 	}
// 	u, err := c.svc.Login(ctx.Request.Context(), req.Email, req.Password)
// 	if err == service.ErrInvalidUserOrPassword {
// 		ctx.String(http.StatusOK, "用户名或者密码不正确，请重试")
// 		return
// 	}
// 	err = c.SetLoginToken(ctx, u.Id)
// 	if err != nil {
// 		ctx.String(http.StatusOK, "系统异常")
// 		return
// 	}
// 	ctx.String(http.StatusOK, "登录成功")
// }

// func (c *UserHandler) Logout(ctx *gin.Context) {
// 	err := c.ClearToken(ctx)
// 	if err != nil {
// 		ctx.JSON(http.StatusOK, Result{
// 			Msg: "系统错误",
// 		})
// 		return
// 	}
// 	ctx.JSON(http.StatusOK, Result{
// 		Msg: "OK",
// 	})
// }

// // Edit 用户编译信息
// func (c *UserHandler) Edit(ctx *gin.Context) {
// 	type Req struct {
// 		// 注意，其它字段，尤其是密码、邮箱和手机，
// 		// 修改都要通过别的手段
// 		// 邮箱和手机都要验证
// 		// 密码更加不用多说了
// 		Nickname string `json:"nickname"`
// 		// 2023-01-01
// 		Birthday string `json:"birthday"`
// 		AboutMe  string `json:"aboutMe"`
// 	}

// 	var req Req
// 	if err := ctx.Bind(&req); err != nil {
// 		return
// 	}
// 	// 你可以尝试在这里校验。
// 	// 比如说你可以要求 Nickname 必须不为空
// 	// 校验规则取决于产品经理
// 	if req.Nickname == "" {
// 		ctx.JSON(http.StatusOK, Result{Code: 4, Msg: "昵称不能为空"})
// 		return
// 	}

// 	if len(req.AboutMe) > 1024 {
// 		ctx.JSON(http.StatusOK, Result{Code: 4, Msg: "关于我过长"})
// 		return
// 	}
// 	birthday, err := time.Parse(time.DateOnly, req.Birthday)
// 	if err != nil {
// 		// 也就是说，我们其实并没有直接校验具体的格式
// 		// 而是如果你能转化过来，那就说明没问题
// 		ctx.JSON(http.StatusOK, Result{Code: 4, Msg: "日期格式不对"})
// 		return
// 	}

// 	uc := ctx.MustGet("user").(ijwt.UserClaims)
// 	err = c.svc.UpdateNonSensitiveInfo(ctx, domain.User{
// 		Id:       uc.Id,
// 		Nickname: req.Nickname,
// 		AboutMe:  req.AboutMe,
// 		Birthday: birthday,
// 	})
// 	if err != nil {
// 		ctx.JSON(http.StatusOK, Result{Code: 5, Msg: "系统错误"})
// 		return
// 	}
// 	ctx.JSON(http.StatusOK, Result{Msg: "OK"})
// }

// // ProfileJWT 用户详情, JWT 版本
// func (c *UserHandler) ProfileJWT(ctx *gin.Context) {
// 	type Profile struct {
// 		Email    string
// 		Phone    string
// 		Nickname string
// 		Birthday string
// 		AboutMe  string
// 	}
// 	uc := ctx.MustGet("user").(ijwt.UserClaims)
// 	u, err := c.svc.Profile(ctx, uc.Id)
// 	if err != nil {
// 		// 按照道理来说，这边 id 对应的数据肯定存在，所以要是没找到，
// 		// 那就说明是系统出了问题。
// 		ctx.String(http.StatusOK, "系统错误")
// 		return
// 	}
// 	ctx.JSON(http.StatusOK, Profile{
// 		Email:    u.Email,
// 		Phone:    u.Phone,
// 		Nickname: u.Nickname,
// 		Birthday: u.Birthday.Format(time.DateOnly),
// 		AboutMe:  u.AboutMe,
// 	})
// }
