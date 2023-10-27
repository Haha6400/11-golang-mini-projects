package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Haha6400/11-golang-mini-projects/go-bookstore/pkg/models"
	"github.com/Haha6400/11-golang-mini-projects/go-bookstore/pkg/utils"
	"github.com/gorilla/mux"
)

var NewBook models.Book

func GetBook(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBooks()
	res, _ := json.Marshal(newBooks)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) //mux.Vars(r) returns a map include variable from URL
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0) //Transform bookId from string to int64
	if err != nil {
		fmt.Printf("error while parsing")
	}
	bookDetails, _ := models.GetBookById(ID)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	newBook := &models.Book{}
	utils.ParseBody(r, newBook)
	b := newBook.CreateBook()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) //mux.Vars(r) returns a map include variable from URL
	bookId := vars["bookId"]
	ID, err := strcov.ParseInt(bookId, 0, 0) //Transform bookId from string to int64
	if err != nil {
		fmt.Printf("error while parsing")
	}
	book := models.DeleteBook(ID)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	var updaeBook = &models.Book
	utils.ParseBody(r, updateBook)
	vars := mux.Vars(r) //mux.Vars(r) returns a map include variable from URL
	bookId := vars["bookId"]
	ID, err := strcov.ParseInt(bookId, 0, 0) //Transform bookId from string to int64
	if err != nil {
		fmt.Printf("error while parsing")
	}
	bookDetails, db := models.GetBookById(ID)
	if updaeBook.Name != "" {
		bookDetails.Name = updaeBook.Name
	}
	if updaeBook.Author != "" {
		bookDetails.Author = updaeBook.Author
	}
	if updaeBook.Publication != "" {
		bookDetails.Publication = updaeBook.Publication
	}
	db.Save(&bookDetails)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
