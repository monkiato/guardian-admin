package routes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"monkiato/guardian-admin/internal/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

//Admin objects required by the Admin routes
type Admin struct {
	db     *gorm.DB
	logger *log.Logger
}

//NewAdmin Admin constructor
func NewAdmin(_db *gorm.DB, _logger *log.Logger) *Admin {
	return &Admin{
		db:     _db,
		logger: _logger,
	}
}

//AddRoutes add all routers related to the users
func (a *Admin) AddRoutes(router *mux.Router) {
	a.logger.Print("adding admin routes...")
	router.HandleFunc("/users", a.getUsersHandler).Methods(http.MethodGet)
	router.HandleFunc("/users/{id}", a.getUserHandler).Methods(http.MethodGet)
	// router.HandleFunc("/users/", a.createUserHandler).Methods(http.MethodPost)
	router.HandleFunc("/users/{id}", a.updateUserHandler).Methods(http.MethodPut)
	// router.HandleFunc("/users/{id}", a.deleteUserHandler).Methods(http.MethodDelete)
}

func (a *Admin) getUsersHandler(w http.ResponseWriter, req *http.Request) {
	// get all users
	users, err := models.GetUsers(a.db)
	if err != nil {
		fmt.Println("Error fetching users", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("unable to fetch users"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Total-Count", strconv.FormatInt(int64(len(users)), 10))
	w.WriteHeader(http.StatusOK)
	responseBody, _ := json.Marshal(users)
	w.Write(responseBody)
	return
}

func (a *Admin) getUserHandler(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	userID, _ := vars["id"]

	// get a single user
	user, err := models.GetUser(a.db, userID)
	if err != nil {
		fmt.Println("Error fetching user "+userID, err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("unable to fetch user"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	responseBody, _ := json.Marshal(user)
	w.Write(responseBody)
	return
}

func (a *Admin) updateUserHandler(w http.ResponseWriter, req *http.Request) {
	var user models.User

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Println("Error reading body data", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("unable to update user"))
		return
	}
	if err := json.Unmarshal(body, &user); err != nil {
		fmt.Println("Error parsing user data", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("unable to update user"))
		return
	}

	if err := models.UpdateUser(a.db, &user); err != nil {
		fmt.Println("Error updating user", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("unable to update user"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(body)
	return
}
