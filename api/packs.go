package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	pc "github.com/sarataha/packs-calculator/pack_calculator"
)

func PacksHandler(w http.ResponseWriter, r *http.Request) {
	itemsStr := r.URL.Query().Get("items")
	items, err := strconv.Atoi(itemsStr)
	if err != nil {
		http.Error(w, "Invalid input: items must be a valid integer", http.StatusBadRequest)
		return
	}

	packs := pc.CalculatePacks(items)

	packsResponse := make([]map[string]int, 0)
	for _, pack := range packs {
		packsResponse = append(packsResponse, map[string]int{"Size": pack.Size, "Quantity": pack.Quantity})
	}

	response := struct {
		Items int
		Packs []map[string]int
	}{Items: items, Packs: packsResponse}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
