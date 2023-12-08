package controller


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
