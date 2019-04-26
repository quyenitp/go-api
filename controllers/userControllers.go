package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"../models"

	"github.com/gorilla/mux"
)

var Users []models.User

//var AccountsInUser = []models.Account{}

//GetUserByID is func get user info by id
func GetUserByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	for _, item := range Users {
		userID, err := strconv.Atoi(params["id"])
		if err != nil {
			log.Fatalln(err)
		}
		if item.ID == userID {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

}

func GetAccountByUserID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	accs := []models.Account{}
	for _, item := range Accounts {
		userID, err := strconv.Atoi(params["id"])
		if err != nil {
			log.Fatalln(err)
		}
		if item.UserID == userID {
			accs = append(accs, item)
			fmt.Println(item.ID)
		}
	}
	json.NewEncoder(w).Encode(accs)
	return
}
