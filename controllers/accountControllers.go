package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"go-api/models"

	"github.com/gorilla/mux"
)

//Accounts will store sample data for Accounts
var Accounts []models.Account

//GetAccountByID is func get Account info by id
func GetAccountByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	for _, item := range Accounts {
		ID, err := strconv.Atoi(params["id"])
		if err != nil {
			log.Fatalln(err)
		}
		if item.ID == ID {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

}
