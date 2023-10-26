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
	db.AutoMigrate(&Book{}) //Auto update Book db table 
}

func (b *Book) CreateBook() *Book{
	db.NewRecord(b) //return boolean if b already exists
	db.Create(b)
	return b
}

func GetAllBooks() []Book{
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(Id int64) (*Book, *gorm.DB){
	var getBook Book
	db := db.Where("ID=?", Id).Find(&getBook)
	return &getBook, db
}

func DeleteBook(Id int64) Book{
	var book Book
	db.Where("ID=?", Id).Delete(book)
	return book
}