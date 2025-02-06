package routers

import (
	"github.com/M1keTrike/API_longPShortP_GO/src/tables/infraestructure/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterTableRoutes(router *gin.Engine, controller *controllers.TableController) {
	router.GET("/longpoll/tables/:table_name", controller.Execute)
}
