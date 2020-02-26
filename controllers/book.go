package controllers

import (
	"books-list/models"
	"books-list/repository/book"
	"database/sql"
	"net/http"
	"log"
	"books-list/utils"
	"github.com/gorilla/mux"
	"strconv"
	"encoding/json"
)
type Controller struct {}


var books [] models.Book

func (c Controller) GetBooks(db *sql.DB) http.HandlerFunc {
	return  func (w http.ResponseWriter, r *http.Request) {
		log.Println("Get  books is called")
		var book models.Book
		var error models.Error
		books =  []models.Book{}
		bookRepo :=  bookRepository.BookRepository{}
		books, err := bookRepo.GetBooks(db, book, books)

		if err != nil {
			error.Message = "Server error"
			utils.SendError(w, http.StatusInternalServerError, error)
			return
		}
		log.Println("books", books)
		w.Header().Set("Content-Type", "application/json")
		utils.SendSuccess(w,books )	
	}
}

func  (c Controller) GetBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Get  book is called")
		var book models.Book
		var error models.Error
		params := mux.Vars(r)
		bookRepo := bookRepository.BookRepository{}
		
		id, _ := strconv.Atoi(params["id"])
		book, err := bookRepo.GetBook(db, book,  id)
		if err != nil {
			log.Println("error:", err)
			if err == sql.ErrNoRows {
				error.Message = "Not found"
				utils.SendError(w,  http.StatusNotFound, error)
				return
			} else {
				error.Message = "Server error"
				utils.SendError(w, http.StatusInternalServerError, error)
				return
			}

		}
		log.Println("book", book)
	    w.Header().Set("Content-Type", "application/json")
		utils.SendSuccess(w,book)	
	}
}

func  (c Controller) AddBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Add book is called")
		var book models.Book
		var bookID int
		var error models.Error
		json.NewDecoder(r.Body).Decode(&book)
		// data checker
		if book.Author == "" || book.Title == "" || book.Year == "" {
			error.Message = "Enter missing fields"
			utils.SendError(w, http.StatusBadRequest, error)
			log.Println("error:", error)
			return
		}
		bookRepo := bookRepository.BookRepository{}
		bookID, err := bookRepo.AddBook(db, book)
		if err != nil {
			error.Message = "Server error"
			utils.SendError(w, http.StatusInternalServerError, error)
			log.Println("error:", error)
			return
		}
		w.Header().Set("Content-Type", "text/plain")
		utils.SendSuccess(w,bookID)	
	}
}

func (c Controller) UpdateBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Update book is called")
		var book models.Book
		var error models.Error
		json.NewDecoder(r.Body).Decode(&book)
		// data checker
		if book.ID==0 || book.Author == "" || book.Title == "" || book.Year == "" {
			error.Message = "All fields are required"
			utils.SendError(w, http.StatusBadRequest, error)
			log.Println("error:", error)
			return
		}
		bookRepo := bookRepository.BookRepository{}
		rowsUpdated, err := bookRepo.UpdateBook(db, book)
		if err != nil {
			error.Message = "Server error"
			utils.SendError(w, http.StatusInternalServerError, error)
			return
		}
		w.Header().Set("Content-Type", "text/plain")
		utils.SendSuccess(w, rowsUpdated)	
	}
}

func (c Controller) RemoveBook(db *sql.DB) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		var error models.Error
		log.Println("Remove book is called")
		params := mux.Vars(r)
		bookRepo := bookRepository.BookRepository{}
		id, _ := strconv.Atoi(params["id"])
		
		rowsDeleted, err := bookRepo.RemoveBook(db, id)
		if err != nil {
			error.Message = "Server error"
			utils.SendError(w, http.StatusInternalServerError, error)
			return
		}
		
		if rowsDeleted == 0 {
			error.Message = "Not found"
			utils.SendError(w, http.StatusNotFound, error)
			return
		}
		w.Header().Set("Content-Type", "text/plain")
		utils.SendSuccess(w, rowsDeleted)	
	}
}