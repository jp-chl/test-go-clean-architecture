package domain

import "context"

type SignUpRequest struct {
	Name     string `form:"name"`
	Login    string `form:"login"`
	Password string `form:"password"`
}

type SignUpResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type SignUpUseCase interface {
	Create(user *User) error
	GetUserByLogin(c context.Context, login string) (User, error)
	CreateAccessToken(user *User, secret string, expiry int) (accessToken string, err error)
	CreateRefreshToken(user *User, secret string, expiry int) (refreshToken string, err error)
}
