package models

import (
	"github.com/Kleist47/go-bookstore/pkg/config" //- возможно следует переложить в публичный репозиторий
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:""json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
}
