package application

import (
	"github.com/M1keTrike/API_longPShortP_GO/src/prices/infraestructure/tech"
)

type PriceChangeUseCase struct {
	repo *tech.PriceChangesRepository
}

func NewPriceChangeUseCase(repo *tech.PriceChangesRepository) *PriceChangeUseCase {
	return &PriceChangeUseCase{repo}
}

func (uc *PriceChangeUseCase) Execute() {
	go uc.repo.MonitorPriceChanges()
}
