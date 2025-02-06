package dependencies

import (
	"os"
	"sync"

	"github.com/M1keTrike/API_longPShortP_GO/src/tables/application"
	"github.com/M1keTrike/API_longPShortP_GO/src/tables/domain/entities"
	"github.com/gin-gonic/gin"

	"github.com/M1keTrike/API_longPShortP_GO/src/tables/infraestructure/controllers"
	"github.com/M1keTrike/API_longPShortP_GO/src/tables/infraestructure/routers"
	"github.com/M1keTrike/API_longPShortP_GO/src/tables/infraestructure/tech"
)

type TablesDependencies struct {
	mutex        *sync.Mutex
	tableChanges *[]entities.TableChange
}

func NewTablesDependencies() *TablesDependencies {
	return &TablesDependencies{
		mutex:        &sync.Mutex{},
		tableChanges: &[]entities.TableChange{},
	}
}

func (d *TablesDependencies) Execute(r *gin.Engine) {
	ep := os.Getenv("MONITOR_API_URL")

	repo := tech.NewTableChangesRepository(ep, d.mutex, d.tableChanges)

	tableChangeUseCase := application.NewTableChangeUseCase(repo)
	tableUnvisitedUseCase := application.NewTableUnvisistedUseCase()

	controller := controllers.NewTableController(tableUnvisitedUseCase, d.mutex, d.tableChanges)

	routers.RegisterTableRoutes(r, controller)

	tableChangeUseCase.Execute()
}
