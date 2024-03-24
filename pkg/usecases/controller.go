package usecases

import "github.com/rodblg/TinURL-Go/pkg/http/apiTinyUrl"

type Usecases struct {
	Storage *apiTinyUrl.Storage
}

func NewUsecasesController() *Usecases {
	return &Usecases{}
}
