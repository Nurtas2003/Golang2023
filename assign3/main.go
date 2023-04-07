package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
)

type Book struct {
	gorm.Model
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Cost        float32 `json:"cost"`
}

var db *gorm.DB

func main() {
	// Connect to the database
	dsn := "postgres:t$hZw!Kz@tcp(localhost:5432)/assign3?charset=utf8mb4&parseTime=True&loc=Local"

	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	// Auto migrate the database schema
	db.AutoMigrate(&Book{})

	// Initialize Gin router
	r := gin.Default()

	// Define REST endpoints
	r.GET("/books", getBooks)
	r.GET("/books/:id", getBookByID)
	r.POST("/books", addBook)
	r.PUT("/books/:id", updateBookByID)
	r.DELETE("/books/:id", deleteBookByID)
	r.GET("/search", searchBooks)
	r.GET("/sorted-books", getSortedBooks)

	// Start the server
	if err := http.ListenAndServe(":8080", r); err != nil {
		panic("Failed to start server")
	}
}

// Get all books
func getBooks(c *gin.Context) {
	var books []Book
	db.Find(&books)
	c.JSON(http.StatusOK, books)
}

// Get a book by ID
func getBookByID(c *gin.Context) {
	var book Book
	if err := db.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}
	c.JSON(http.StatusOK, book)
}

// Add a new book
func addBook(c *gin.Context) {
	var book Book
	if err := c.BindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Create(&book)
	c.JSON(http.StatusCreated, book)
}

// Update a book by ID
func updateBookByID(c *gin.Context) {
	var book Book
	if err := db.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}
	if err := c.BindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Save(&book)
	c.JSON(http.StatusOK, book)
}

// Delete a book by ID
func deleteBookByID(c *gin.Context) {
	var book Book
	if err := db.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}
	db.Delete(&book)
	c.Status(http.StatusNoContent)
}

// Search for books by title
func searchBooks(c *gin.Context) {
	var books []Book
	if err := db.Where("title LIKE ?", "%"+c.Query("title")+"%").Find(&books).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to search books"})
		return
	}
	c.JSON(http.StatusOK, books)
}

// Get a sorted list of books ordered by cost in descending or ascending order
func getSortedBooks(c *gin.Context) {
	var books []Book
	order := c.Query("order")
	if order == "desc" {
		db.Order("cost desc").Find(&books)
	} else {
		db.Order("cost asc").Find(&books)
	}
	c.JSON(http.StatusOK, books)
}
