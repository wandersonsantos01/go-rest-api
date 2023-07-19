package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/wandersonsantos01/go-rest-api/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home")
}

func AllCars(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.Cars)
}
