package controllers

import (
	"net/http"
	"sync"

	"github.com/M1keTrike/API_longPShortP_GO/src/tables/application"
	"github.com/M1keTrike/API_longPShortP_GO/src/tables/domain/entities"
	"github.com/gin-gonic/gin"
)

type TableController struct {
	unvisitedUseCase *application.TableUnvisitedUseCase
	mutex            *sync.Mutex
	changes          *[]entities.TableChange
}

func NewTableController(unvisitedUseCase *application.TableUnvisitedUseCase, mutex *sync.Mutex, changes *[]entities.TableChange) *TableController {
	return &TableController{
		unvisitedUseCase: unvisitedUseCase,
		mutex:            mutex,
		changes:          changes,
	}
}

func (tc *TableController) Execute(c *gin.Context) {
	tableName := c.Param("table_name")

	changes := tc.unvisitedUseCase.Execute(tableName, tc.mutex, tc.changes)
	if len(changes) > 0 {
		c.JSON(http.StatusOK, changes)
	} else {
		c.Status(http.StatusNoContent) 
	}
}
