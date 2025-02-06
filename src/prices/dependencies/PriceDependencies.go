package dependencies

import (
	"os"
	"sync"

	"github.com/M1keTrike/API_longPShortP_GO/src/prices/application"
	"github.com/M1keTrike/API_longPShortP_GO/src/prices/domain/entities"
	"github.com/M1keTrike/API_longPShortP_GO/src/prices/infraestructure/controllers"
	"github.com/M1keTrike/API_longPShortP_GO/src/prices/infraestructure/routers"
	"github.com/M1keTrike/API_longPShortP_GO/src/prices/infraestructure/tech"
	"github.com/gin-gonic/gin"
)

type PricesDependencies struct {
	mutex       *sync.Mutex
	priceChanges *[]entities.PriceChange
}

func NewPricesDependencies() *PricesDependencies {
	return &PricesDependencies{
		mutex:       &sync.Mutex{},
		priceChanges: &[]entities.PriceChange{},
	}
}

func (d *PricesDependencies) Execute(r *gin.Engine) {
	ep := os.Getenv("MONITOR_API_URL_PRICES")

	// Inicializar el repositorio
	repo := tech.NewPriceChangesRepository(ep, d.mutex, d.priceChanges)

	// Inicializar los casos de uso
	priceChangeUseCase := application.NewPriceChangeUseCase(repo)
	priceUnvisitedUseCase := application.NewPriceUnvisitedUseCase()

	// Inicializar el controlador
	controller := controllers.NewPriceController(priceUnvisitedUseCase, d.mutex, d.priceChanges)

	// Registrar rutas
	routers.RegisterPriceRoutes(r, controller)

	// Iniciar el monitoreo de cambios de precios
	priceChangeUseCase.Execute()
}
