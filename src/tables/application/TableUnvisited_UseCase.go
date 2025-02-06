package application

import (
	"sync"
	"time"

	"github.com/M1keTrike/API_longPShortP_GO/src/tables/domain/entities"
)

type TableUnvisitedUseCase struct {}


func NewTableUnvisistedUseCase() *TableUnvisitedUseCase {
	return &TableUnvisitedUseCase{}
}

func (uc *TableUnvisitedUseCase) Execute(tableName string, mutex *sync.Mutex, changes *[]entities.TableChange) []entities.TableChange {
	timeout := time.After(30 * time.Second)

	for {
		mutex.Lock()
		var unvisitedChanges []entities.TableChange
		for i, change := range *changes {
			if change.Table == tableName && !change.Visited {
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
		
			return []entities.TableChange{}
		}
	}
}
