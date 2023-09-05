package repository

import "github.com/jp-chl/test-go-clean-architecture/domain/model"

type RedirectRepository interface {
	Find(code string) (*model.Redirect, error)
	Store(redirect *model.Redirect) error
}
