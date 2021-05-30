package handler

import (
	"encoding/json"
	"net/http"

	"github.com/amol-scalent/poc/rest-api/db"
	"github.com/amol-scalent/poc/rest-api/models"
	"github.com/gorilla/mux"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	DB := db.ConnectToDB() // calling database instance
	w.Header().Set("Content-Type", "application/json")
	var users []models.User
	DB.Find(&users)
	json.NewEncoder(w).Encode(users)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	DB := db.ConnectToDB()
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var user models.User
	DB.First(&user, params["id"])
	json.NewEncoder(w).Encode(user)
}
func CreateUser(w http.ResponseWriter, r *http.Request) {
	DB := db.ConnectToDB()
	w.Header().Set("Content-Type", "application/json")
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)
	DB.Create(&user)
	json.NewEncoder(w).Encode(user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	DB := db.ConnectToDB()
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var user models.User
	DB.First(&user, params["id"])
	json.NewDecoder(r.Body).Decode(&user)
	DB.Save(&user)
	json.NewEncoder(w).Encode(user)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	DB := db.ConnectToDB()
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var user models.User
	DB.Delete(&user, params["id"])
	json.NewEncoder(w).Encode("The USer is Deleted Successfully!")
}
