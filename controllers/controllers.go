package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/wandersonsantos01/go-rest-api/database"
	"github.com/wandersonsantos01/go-rest-api/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home")
}

func AllCars(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-type", "application/json")
	var cars []models.Car
	database.DB.Find(&cars)
	json.NewEncoder(w).Encode(cars)
}

func GetCarById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var car models.Car
	database.DB.First(&car, id)

	json.NewEncoder(w).Encode(car)
}

func CreateCar(w http.ResponseWriter, r *http.Request) {
	var car models.Car
	json.NewDecoder(r.Body).Decode(&car)
	database.DB.Create(&car)
	json.NewEncoder(w).Encode(car)
}

func DeleteCar(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var car models.Car
	database.DB.Delete(&car, id)
	json.NewEncoder(w).Encode(car)
}

func UpdateCar(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var car models.Car
	database.DB.First(&car, id)

	json.NewDecoder(r.Body).Decode(&car)
	database.DB.Save(&car)
	json.NewEncoder(w).Encode(car)
}
