package application

import (
	"sync"
	"time"

	"github.com/M1keTrike/API_longPShortP_GO/src/prices/domain/entities"
)

type PriceUnvisitedUseCase struct {}

func NewPriceUnvisitedUseCase() *PriceUnvisitedUseCase {
	return &PriceUnvisitedUseCase{}
}

func (uc *PriceUnvisitedUseCase) Execute(productID int, mutex *sync.Mutex, changes *[]entities.PriceChange) []entities.PriceChange {
	timeout := time.After(30 * time.Second)

	for {
		mutex.Lock()
		var unvisitedChanges []entities.PriceChange
		for i, change := range *changes {
			if change.ProductID == productID && !change.Visited {
				(*changes)[i].Visited = true
				unvisitedChanges = append(unvisitedChanges, change)
			}
		}
		mutex.Unlock()

	
		if len(unvisitedChanges) > 0 {
			return unvisitedChanges
		}

	
		select {
		case <-time.After(1 * time.Second):
		
		case <-timeout:
			
			return []entities.PriceChange{}
		}
	}
}
