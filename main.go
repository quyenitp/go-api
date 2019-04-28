package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/quyenitp/go-api/controllers"
	"github.com/quyenitp/go-api/models"
)

func main() {

	var Users []models.User

	Users = append(Users, models.User{ID: 1, Name: "Jame", AccountIDs: "1,3,5"})
	Users = append(Users, models.User{ID: 2, Name: "Bond", AccountIDs: "4"})
	Users = append(Users, models.User{ID: 3, Name: "Jimmy", AccountIDs: "2"})

	var Accounts []models.Account
	Accounts = append(Accounts, models.Account{ID: 1, UserID: 1, Name: "Tiet Kiem", Balance: 50000})
	Accounts = append(Accounts, models.Account{ID: 2, UserID: 3, Name: "Thanh Toan", Balance: 10000})
	Accounts = append(Accounts, models.Account{ID: 3, UserID: 1, Name: "Vay Tin Dung", Balance: 150000})
	Accounts = append(Accounts, models.Account{ID: 4, UserID: 2, Name: "Thanh Toan", Balance: 70000})
	Accounts = append(Accounts, models.Account{ID: 5, UserID: 1, Name: "Thanh Toan", Balance: 70000})

	controllers.Users = Users

	controllers.Accounts = Accounts

	fmt.Println("Running RESTfull API ... on port 8080")
	router := mux.NewRouter()

	router.HandleFunc("/users/{id}", controllers.GetUserByID).Methods("GET")
	router.HandleFunc("/users/{id}/accounts", controllers.GetAccountByUserID).Methods("GET")
	router.HandleFunc("/accounts/{id}", controllers.GetAccountByID).Methods("GET")
	log.Fatalln(http.ListenAndServe(":8080", router))
}
