package main

import (
	"fmt"
	a "github.com/NurtasSerikkanov/Golang2023/assign3/pkg"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"strconv"
)

var db *gorm.DB

func main() {
	// Load environment variables
	err := godotenv.Load(".env")

	// Create DSN string for connecting to Postgres database
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Almaty",
		os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))

	// Parse port to integer
	_, err = strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		log.Fatalf("Error parsing DB_PORT: %v", err)
	}

	// Connect to Postgres database using gorm
	db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	err = db.AutoMigrate(&a.Book{})

	c := a.Connection{DB: db}

	router := mux.NewRouter()

	router.HandleFunc("/books/", c.GetAllBooks).Methods("GET")
	router.HandleFunc("/books/{id}/", c.GetBookByID).Methods("GET")
	router.HandleFunc("/addBook/", c.AddBook).Methods("POST")
	router.HandleFunc("/updateBook/{id}/", c.UpdateBook).Methods("PUT")
	router.HandleFunc("/deleteBook/{id}/", c.DeleteBookByID).Methods("DELETE")
	router.HandleFunc("/search/{title}/", c.SearchBookByTitle).Methods("GET")
	router.HandleFunc("/sortedBooks/", c.GetSortedBooks).Methods("GET")
	router.HandleFunc("/descSortedBooks/", c.DescGetSortedBooks).Methods("GET")

	http.ListenAndServe(":8080", router)
}
