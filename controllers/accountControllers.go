package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"../models"
	"github.com/gorilla/mux"
)

var Accounts []models.Account

//GetUserByID is func get user info by id
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
