package tech

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/M1keTrike/API_longPShortP_GO/src/prices/domain/entities"
)

type PriceChangesRepository struct {
	Ep      string
	Mutex   *sync.Mutex
	Changes *[]entities.PriceChange
}

func NewPriceChangesRepository(ep string, mutex *sync.Mutex, changes *[]entities.PriceChange) *PriceChangesRepository {
	return &PriceChangesRepository{Ep: ep, Mutex: mutex, Changes: changes}
}

func (r *PriceChangesRepository) MonitorPriceChanges() {
	var previousChanges map[int]entities.PriceChange = make(map[int]entities.PriceChange)

	for {
		resp, err := http.Get(r.Ep)
		if err != nil {
			log.Println("Error al obtener cambios en los precios:", err)
			time.Sleep(5 * time.Second)
			continue
		}

		if resp.StatusCode == http.StatusOK {
			var response struct {
				PriceChanges []entities.PriceChange `json:"price_changes"`
			}

			if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
				log.Println("Error al decodificar respuesta JSON:", err)
				resp.Body.Close()
				continue
			}

			resp.Body.Close()

			r.Mutex.Lock()
			newChanges := make([]entities.PriceChange, 0)

			for _, change := range response.PriceChanges {
				
				if prevChange, exists := previousChanges[change.ProductID]; !exists || prevChange.NewPrice != change.NewPrice {
					change.Visited = false
					newChanges = append(newChanges, change)
					previousChanges[change.ProductID] = change
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
