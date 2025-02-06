package application

import (


	"github.com/M1keTrike/API_longPShortP_GO/src/tables/infraestructure/tech"
)

type TableChangeUseCase struct {
	repo *tech.TableChangesRepository
}


func NewTableChangeUseCase(repo *tech.TableChangesRepository) *TableChangeUseCase {
	return &TableChangeUseCase{repo}
}


func (uc *TableChangeUseCase) Execute() {
	go uc.repo.MonitorTableChanges()
}


