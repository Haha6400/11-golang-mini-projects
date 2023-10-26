package models

import (
	"github.com/Haha6400/11-golang-mini-projects/go-bookstore/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm: ""json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.getDB()
	db.AutoMigrate(&Book{})
}
