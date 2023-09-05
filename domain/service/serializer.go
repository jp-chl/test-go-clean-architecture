package service

import "github.com/jp-chl/test-go-clean-architecture/domain/model"

type RedirectSerializer interface {
	Decode(input []byte) (*model.Redirect, error)
	Encode(input *model.Redirect) ([]byte, error)
}
