package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/wandersonsantos01/go-rest-api/controllers"
	"github.com/wandersonsantos01/go-rest-api/routes/middleware"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware)
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/cars", controllers.AllCars).Methods("Get")
	r.HandleFunc("/api/cars/{id}", controllers.GetCarById).Methods("Get")
	r.HandleFunc("/api/cars", controllers.CreateCar).Methods("Post")
	r.HandleFunc("/api/cars/{id}", controllers.DeleteCar).Methods("Delete")
	r.HandleFunc("/api/cars/{id}", controllers.UpdateCar).Methods("Put")
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}
