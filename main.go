package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Mohammad-Hakemi22/mongoAPI/router"
)

func main() {
	fmt.Println("movies whatched list API")
	fmt.Println("Starting server ...")
	r := router.Router()
	log.Fatal(http.ListenAndServe(":8000", r))

}