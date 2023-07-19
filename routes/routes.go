package routes

import (
	"log"
	"net/http"

	"github.com/wandersonsantos01/go-rest-api/controllers"
)

func HandleRequest() {
	http.HandleFunc("/", controllers.Home)
	http.HandleFunc("/api/cars", controllers.AllCars)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
