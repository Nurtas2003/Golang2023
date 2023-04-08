package main

import (
	"fmt"
	a "github.com/NurtasSerikkanov/Golang2023/assign3/pkg"
	"github.com/gin-gonic/gin"
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
	r := gin.Default()

	// Define REST endpoints
	r.GET("/books", a.GetBooks)
	r.GET("/books/:id", a.GetBookByID)
	r.POST("/books", a.AddBook)
	r.PUT("/books/:id", a.UpdateBookByID)
	r.DELETE("/books/:id", a.DeleteBookByID)
	r.GET("/search", a.SearchBookByName)
	r.GET("/sorted-books", a.GetSortedBooks)

	//db.AutoMigrate(&.Book{})
	//router := mux.NewRouter()
	//router.HandleFunc("/", homePage).Methods("GET")
	//router.HandleFunc("/books/", a.GetBooks).Methods("GET")
	//router.HandleFunc("/books/sort/", a.GetSortedBooks).Methods("GET")
	//router.HandleFunc("/books/{id}/", a.getBookById).Methods("GET")
	//router.HandleFunc("/books/add/", a.AddBook).Methods("POST")
	//router.HandleFunc("/books/{id}/update/", a.UpdateBook).Methods("PUT")
	//router.HandleFunc("/books/{id}/delete/", a.DeleteBook).Methods("DELETE")
	//router.HandleFunc("/books/", a.SearchBookByName).Methods("GET")
	// Start the server
	fmt.Println("Server at 8080")
	errr := http.ListenAndServe(":7575", r)
	if errr != nil {
		fmt.Println(errr)
	}
}
