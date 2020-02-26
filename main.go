package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	 "books-list/models"
	 "books-list/driver"
	 "books-list/controllers"
	"database/sql"
	"github.com/subosito/gotenv"
	"github.com/gorilla/handlers"
	"os"
)

var  books [] models.Book;
var db *sql.DB;
func init() {
	gotenv.Load()
}
func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
func main() {
	  db = driver.ConnectDB()
	  controller := controllers.Controller{}

	router :=  mux.NewRouter()

	// set up the router
	router.HandleFunc("/books", controller.GetBooks(db)).Methods("GET")
	router.HandleFunc("/books/{id}", controller.GetBook(db)).Methods("GET")
	router.HandleFunc("/books",  controller.AddBook(db)).Methods("POST")
	router.HandleFunc("/books", controller.UpdateBook(db)).Methods("PUT")
	router.HandleFunc("/books/{id}", controller.RemoveBook(db)).Methods("DELETE")
	log.Println("Server is running at port 8000")
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	originsOk := handlers.AllowedOrigins([]string{os.Getenv("ORIGIN_ALLOWED")})
    methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS","DELETE"})
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(headersOk, originsOk, methodsOk)(router)))

}