package handlers

import "github.com/jinzhu/gorm"

type Handler struct {
	Database *gorm.DB
}
