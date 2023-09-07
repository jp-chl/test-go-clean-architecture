package serializer

import (
	"encoding/json"
	"errors"

	"github.com/jp-chl/test-go-clean-architecture/domain/model"
)

type Redirect struct{}

func (r *Redirect) Decode(input []byte) (*model.Redirect, error) {
	redirect := &model.Redirect{}
	if err := json.Unmarshal(input, redirect); err != nil {
		return nil, errors.New("unable to Unmarshal")
	}
	return redirect, nil
}

func (r *Redirect) Encode(input *model.Redirect) ([]byte, error) {
	message, err := json.Marshal(input)
	if err != nil {
		return nil, errors.New("unable to Marshall")
	}
	return message, nil
}
