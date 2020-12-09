package models

import (
	"errors"
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

//User DB schema
type User struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`

	Avatar        string `gorm:"type:varchar(255);null" json:"avatar"`
	Username      string `gorm:"type:varchar(100);unique_index;not null" json:"username"`
	Name          string `gorm:"type:varchar(100);not null" json:"name"`
	Lastname      string `gorm:"type:varchar(100);not null" json:"lastname"`
	Password      string `gorm:"type:varchar(100);not null" json:"password"`
	Email         string `gorm:"type:varchar(100);unique;not null" json:"email"`
	Token         string `gorm:"type:varchar(255);unique;not null" json:"token"`
	ApprovalToken string `gorm:"type:varchar(255);not null" json:"approval_token"`
	Approved      bool   `gorm:"type:boolean" json:"approved"`
}

//GetUsers get all user records
func GetUsers(db *gorm.DB) ([]User, error) {
	var users []User

	errs := db.Find(&users).GetErrors()
	if len(errs) > 0 {
		for _, err := range errs {
			fmt.Errorf(err.Error())
		}
		return nil, errors.New("record not found")
	}
	return users, nil
}

//CreateUser create a new user
func CreateUser(db *gorm.DB, user *User) error {
	return db.Create(&user).Error
}

//GetUser get a single user record
func GetUser(db *gorm.DB, userID string) (*User, error) {
	var user User

	if db.Where("id = ?", userID).First(&user).RecordNotFound() {
		return nil, errors.New("record not found")
	}
	return &user, nil
}

//UpdateUser update an existing user
func UpdateUser(db *gorm.DB, user *User) error {
	return db.Save(&user).Error
}
