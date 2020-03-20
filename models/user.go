package models

import (
	_ "encoding/json"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	gorm.Model
	Username string `json:"username"`
	Secret   string `json:"secret"`
}

type Users []User
