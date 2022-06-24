package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Mohammad-Hakemi22/mongoAPI/connections"
)

func getAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/x-www-from-urlencode")
	allMovies := connections.GetAllMovies_helper()
	json.NewEncoder(w).Encode(allMovies)
}
