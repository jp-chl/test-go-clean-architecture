package domain

import "context"

type LoginRequest struct {
	Login    string `form:"login"`
	Password string `form:"password"`
}

type LoginResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type LoginUseCase interface {
	GetUserByLogin(c context.Context, login string) (User, error)
	CreateAccessToken(user *User, secret string, expiry int) (accessToken string, err error)
	CreateRefreshToken(user *User, secret string, expiry int) (refreshToken string, err error)
}
