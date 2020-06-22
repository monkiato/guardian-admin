package models

import (
	"errors"
	"fmt"

	"github.com/jinzhu/gorm"
)

//User DB schema
type User struct {
	gorm.Model

	Username      string `gorm:"type:varchar(100);unique_index;not null"`
	Name          string `gorm:"type:varchar(100);not null"`
	Lastname      string `gorm:"type:varchar(100);not null"`
	Password      string `gorm:"type:varchar(100);not null"`
	Email         string `gorm:"type:varchar(100);unique;not null"`
	Token         string `gorm:"type:varchar(255);unique;not null"`
	ApprovalToken string `gorm:"type:varchar(255);not null"`
	Approved      bool   `gorm:"type:boolean"`
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
