package service

import "github.com/jp-chl/test-go-clean-architecture/domain/model"

type RedirectService interface {
	Find(code string) (*model.Redirect, error)
	Store(redirect *model.Redirect) error
}
