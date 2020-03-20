package models

import (
	_ "encoding/json"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Credentials struct {
	username string `json:"username"`
	Password string `json:"password"`
}
