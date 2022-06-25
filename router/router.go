package router

import (
	"github.com/Mohammad-Hakemi22/mongoAPI/controllers"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/movies", controllers.GetAllMovies).Methods("GET")
	router.HandleFunc("/api/movie", controllers.CreateMovie).Methods("POST")
	router.HandleFunc("/api/movie/{id}", controllers.MarkAsWhatched).Methods("PUT")
	router.HandleFunc("/api/movie/{id}", controllers.DeleteMovie).Methods("DELETE")
	router.HandleFunc("/api/delmovies", controllers.DeleteAllMovies).Methods("DELETE")
	return router
}