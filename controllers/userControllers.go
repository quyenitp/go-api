package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/quyenitp/go-api/models"
)

//Users will store sample value from main func
var Users []models.User

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

//GetAccountByUserID is func get All account info by UserID
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
		}
	}
	json.NewEncoder(w).Encode(accs)
	return
}
