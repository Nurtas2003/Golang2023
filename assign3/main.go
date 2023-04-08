package main

import (
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"gorm.io/gorm"
	"net/http"
)

var db *gorm.DB

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello")
}

func main() {
	// Initialize Gin router
	//r := gin.Default()

	// Define REST endpoints
	//r.GET("/books", getBooks)
	//r.GET("/books/:id", getBookByID)
	//r.POST("/books", addBook)
	//r.PUT("/books/:id", updateBookByID)
	//r.DELETE("/books/:id", deleteBookByID)
	//r.GET("/search", searchBooks)
	//r.GET("/sorted-books", getSortedBooks)

	//db.AutoMigrate(&pkg.Book{})
	router := mux.NewRouter()

	router.HandleFunc("/", homePage).Methods("GET")
	router.HandleFunc("/books", controller.getBooks).Methods("GET")
	router.HandleFunc("/books/", controller.getSortedBooks).Methods("GET")
	router.HandleFunc("/books/{id}", controller.getBooksById).Methods("GET")
	router.HandleFunc("/books", controller.addBook).Methods("POST")
	router.HandleFunc("/books/{id}", controller.updateBook).Methods("PUT")
	// Start the server
	fmt.Println("Server at 8080")
	errr := http.ListenAndServe(":7575", router)
	if errr != nil {
		fmt.Println(errr)
	}
}
