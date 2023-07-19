package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/wandersonsantos01/go-rest-api/controllers"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/cars", controllers.AllCars).Methods("Get")
	r.HandleFunc("/api/cars/{id}", controllers.GetCarById).Methods("Get")
	log.Fatal(http.ListenAndServe(":8000", r))
}
