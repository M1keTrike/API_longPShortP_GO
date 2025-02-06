package tech

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/M1keTrike/API_longPShortP_GO/src/tables/domain/entities"
)


type TableChangesRepository struct {
	Ep     string
	Mutex  *sync.Mutex
	Changes *[]entities.TableChange
}

func NewTableChangesRepository(ep string, mutex *sync.Mutex, changes *[]entities.TableChange) *TableChangesRepository {
	return &TableChangesRepository{Ep: ep, Mutex: mutex, Changes: changes}
}

func (r *TableChangesRepository) MonitorTableChanges() {
	var previousChanges map[string]entities.TableChange = make(map[string]entities.TableChange)

	for {
		resp, err := http.Get(r.Ep)
		if err != nil {
			log.Println("Error al obtener cambios en tablas:", err)
			time.Sleep(5 * time.Second)
			continue
		}

		if resp.StatusCode == http.StatusOK {
			var response struct {
				TableChanges []entities.TableChange `json:"table_changes"`
			}

			if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
				log.Println("Error al decodificar respuesta JSON:", err)
				resp.Body.Close()
				continue
			}

			resp.Body.Close()

			r.Mutex.Lock()
			newChanges := make([]entities.TableChange, 0)

			for _, change := range response.TableChanges {
				key := change.Table + "-" + change.Action + "-" + change.EventTime 

				// Si el cambio no existía antes, es nuevo
				if _, exists := previousChanges[key]; !exists {
					change.Visited = false
					newChanges = append(newChanges, change)
					previousChanges[key] = change
				}
			}

			*r.Changes = append(*r.Changes, newChanges...)

			r.Mutex.Unlock()
		} else {
			resp.Body.Close()
		}

		time.Sleep(5 * time.Second)
	}
}


