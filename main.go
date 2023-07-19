package main

import (
	"fmt"

	"github.com/wandersonsantos01/go-rest-api/models"
	"github.com/wandersonsantos01/go-rest-api/routes"
)

const _default_port = "8000"

func main() {

	models.Cars = []models.Car{
		{Id: 1, Name: "Fusca", History: "History of fusca"},
		{Id: 2, Name: "Chevete", History: "History of chevete"},
	}

	fmt.Println("Server started on port", "8000")
	routes.HandleRequest()

}
