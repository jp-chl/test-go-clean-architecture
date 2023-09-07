package redirect

import (
	// "errors"
	"time"

	"github.com/jp-chl/test-go-clean-architecture/domain/model"
	"github.com/jp-chl/test-go-clean-architecture/domain/repository"
	"github.com/jp-chl/test-go-clean-architecture/domain/service"
)

// var (
// 	ErrRedirectNotFound = errors.New("redirect Not Found")
// 	ErrRedirectInvalid  = errors.New("redirect Invalid")
// )

type redirectService struct {
	redirectRepository repository.RedirectRepository
}

func NewRedirectService(redirectRepository repository.RedirectRepository) service.RedirectService {
	return &redirectService{
		redirectRepository: redirectRepository,
	}
}

func (r *redirectService) Find(code string) (*model.Redirect, error) {
	return r.redirectRepository.Find(code)
}

func (r *redirectService) Store(redirect *model.Redirect) error {
	redirect.Code = "123"
	redirect.CreatedAt = time.Now().UTC().Unix()

	return r.redirectRepository.Store(redirect)
}
