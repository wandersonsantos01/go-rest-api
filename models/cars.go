package models

type Car struct {
	Name    string `json:"name"`
	History string `json:"history"`
}

var Cars []Car
