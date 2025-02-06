package routers

import (
	"github.com/M1keTrike/API_longPShortP_GO/src/prices/infraestructure/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterPriceRoutes(router *gin.Engine, controller *controllers.PriceController) {
	router.GET("/longpoll/prices/:product_id", controller.Execute)
}
