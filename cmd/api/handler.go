package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Inventory struct {
	ID          int32  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type AllInventory struct {
	Data []Inventory `json:"data"`
}

func (app *Config) CreateInventory(w http.ResponseWriter, r *http.Request) {

	allInventory := []Inventory{
		{ID: 1, Name: "shoes", Description: "size 10"},
		{ID: 1, Name: "trouses", Description: "size 20"},
		{ID: 1, Name: "shirts", Description: "medium"},
	}

	data := &AllInventory{Data: allInventory}

	out, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)

	_, err = w.Write(out)
	if err != nil {
		log.Println(err)
	}
}
