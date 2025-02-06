package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	priceDeps "github.com/M1keTrike/API_longPShortP_GO/src/prices/dependencies"
	tableDeps "github.com/M1keTrike/API_longPShortP_GO/src/tables/dependencies"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error cargando el archivo .env:", err)
	}

	r := gin.Default()

	tablesDeps := tableDeps.NewTablesDependencies()
	tablesDeps.Execute(r)

	pricesDeps := priceDeps.NewPricesDependencies()
	pricesDeps.Execute(r)

	log.Println("Servidor corriendo en :8082")
	r.Run(":8082")
}
