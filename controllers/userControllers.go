package controllers

import (
	"encoding/json"
	"errors"

	"github.com/gorilla/mux"
	"go-api/models"
	"log"
	"net/http"
	"strconv"
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

func accountIDExistsInAccountIDs(accountIDs []int, account int) bool {
	for _, accID := range accountIDs {
		if accID == account {
			return true
		}
	}
	return false
}

func getUserByID(userID int) (models.User, error) {
	var user models.User
	for _, u := range Users {
		if u.ID == userID {
			return u, nil
		}
	}
	return user, errors.New("User not found")
}

//GetAccountByUserID is func get All account info by UserID
func GetAccountByUserID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	accs := []models.Account{}
	for _, acc := range Accounts {
		userID, err := strconv.Atoi(params["id"])
		if err != nil {
			log.Println("Cannot parse params Id")
			return
		}

		userInfo, err := getUserByID(userID)
		if err != nil {
			log.Println(err)
			return
		}

		if acc.UserID == userID && accountIDExistsInAccountIDs(userInfo.AccountIDs, acc.ID) == true {
			accs = append(accs, acc)
		}
	}
	json.NewEncoder(w).Encode(accs)
	return
}
