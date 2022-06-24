package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Mohammad-Hakemi22/mongoAPI/connections"
	"github.com/Mohammad-Hakemi22/mongoAPI/models"
	"github.com/gorilla/mux"
)

func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/x-www-from-urlencode")
	allMovies := connections.GetAllMovies_helper()
	err := json.NewEncoder(w).Encode(allMovies)
	connections.CheckError(err)
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/x-www-from-urlencode")
	var movie models.Netflix
	err := json.NewDecoder(r.Body).Decode(&movie)
	connections.CheckError(err)
	connections.InsertOneMovie_helper(movie)
	json.NewEncoder(w).Encode(movie)
}

func MarkAsWhatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/x-www-from-urlencode")
	params := mux.Vars(r)
	connections.UpdateOneMovie_helper(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/x-www-from-urlencode")
	params := mux.Vars(r)
	connections.DeleteOneMovie_helper(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/x-www-from-urlencode")
	count := connections.DeleteAllMovies_helper()
	json.NewEncoder(w).Encode(count)
}