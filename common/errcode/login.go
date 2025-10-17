package errcode

import "gin/app/model"

// TokenData 登录返回的Token信息
type TokenData struct {
	AccessToken        string `json:"accessToken" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
	RefreshToken       string `json:"refreshToken" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
	TokenExpire        int64  `json:"tokenExpire" example:"7200"`
	RefreshTokenExpire int64  `json:"refreshTokenExpire" example:"172800"`
}

// LoginData 登录成功
type LoginData struct {
	Token TokenData  `json:"token"`
	User  model.User `json:"user"`
}

// LoginResponse 登录响应
type LoginResponse struct {
	Code int64     `json:"code" example:"0"`
	Msg  string    `json:"msg" example:"Success"`
	Data LoginData `json:"data"`
}
