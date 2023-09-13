package redirect

import (
	// "errors"

	"math/rand"
	"time"

	"github.com/jp-chl/test-go-clean-architecture/domain/model"
	"github.com/jp-chl/test-go-clean-architecture/domain/repository"
	"github.com/jp-chl/test-go-clean-architecture/domain/service"
)

// var (
// 	ErrRedirectNotFound = errors.New("redirect Not Found")
// 	ErrRedirectInvalid  = errors.New("redirect Invalid")
// )

var (
	seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
)

const (
	RANDOM_CODE_MAX_LENGTH = 8
	RANDOM_CODE_CHARSET    = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

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
	redirect.Code = r.randomCode()
	redirect.CreatedAt = time.Now().UTC().Unix()

	return r.redirectRepository.Store(redirect)
}

func (r *redirectService) randomCode() string {
	result := make([]byte, RANDOM_CODE_MAX_LENGTH)
	for i := range result {
		result[i] = RANDOM_CODE_CHARSET[seededRand.Intn(len(RANDOM_CODE_CHARSET))]
	}
	return string(result)
}
