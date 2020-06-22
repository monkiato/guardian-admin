package routes

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"log"
	"monkiato/guardian-admin/internal/models"
	"net/http"
)

//Admin objects required by the Admin routes
type Admin struct {
	db                   *gorm.DB
	logger               *log.Logger
}

//NewAdmin Admin constructor
func NewAdmin(_db *gorm.DB, _logger *log.Logger) *Admin {
	return &Admin{
		db:                   _db,
		logger:               _logger,
	}
}

//AddRoutes add all routers related to the users
func (a *Admin) AddRoutes(router *mux.Router) {
	a.logger.Print("adding admin routes...")
	router.HandleFunc("/users", a.getUsersHandler).Methods("GET")
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

	w.WriteHeader(http.StatusOK)
	responseBody, _ := json.Marshal(users)
	w.Write(responseBody)
	return
}
