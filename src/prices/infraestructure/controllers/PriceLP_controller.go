package controllers

import (
	"net/http"
	"strconv"
	"sync"

	"github.com/M1keTrike/API_longPShortP_GO/src/prices/application"
	"github.com/M1keTrike/API_longPShortP_GO/src/prices/domain/entities"
	"github.com/gin-gonic/gin"
)

type PriceController struct {
	unvisitedUseCase *application.PriceUnvisitedUseCase
	mutex            *sync.Mutex
	changes          *[]entities.PriceChange
}

func NewPriceController(unvisitedUseCase *application.PriceUnvisitedUseCase, mutex *sync.Mutex, changes *[]entities.PriceChange) *PriceController {
	return &PriceController{
		unvisitedUseCase: unvisitedUseCase,
		mutex:            mutex,
		changes:          changes,
	}
}

func (pc *PriceController) Execute(c *gin.Context) {
	productID, err := strconv.Atoi(c.Param("product_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product_id"})
		return
	}

	changes := pc.unvisitedUseCase.Execute(productID, pc.mutex, pc.changes)
	if len(changes) > 0 {
		c.JSON(http.StatusOK, changes)
	} else {
		c.Status(http.StatusNoContent) 
	}
}
