package pkg

import (
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"gorm.io/gorm"
	"net/http"
)

var db *gorm.DB

//db.AutoMigrate(&Book{})

func GetBooks(c *gin.Context) {
	var books []Book
	db.Find(&books)
	c.JSON(http.StatusOK, books)
}

func GetBookByID(c *gin.Context) {
	var book Book
	if err := db.Where("id=?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}
	c.JSON(http.StatusOK, book)
}

func AddBook(c *gin.Context) {
	var book Book
	if err := c.BindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Create(&book)
	c.JSON(http.StatusCreated, book)
}

func UpdateBookByID(c *gin.Context) {
	var book Book
	if err := db.Where("id=?", c.Param("id")).First(&book).Error; err != nil {
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

func DeleteBookByID(c *gin.Context) {
	var book Book
	if err := db.Where("id=?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}
	db.Delete(&book)
	c.Status(http.StatusNoContent)
}

func SearchBookByName(c *gin.Context) {
	var books []Book
	if err := db.Where("title LIKE ?", "%"+c.Query("title")+"%").Find(&books).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed tosearch books"})
		return
	}
	c.JSON(http.StatusOK, books)
}

func GetSortedBooks(c *gin.Context) {
	var books []Book
	order := c.Query("order")
	if order == "desk" {
		db.Order("cost desc").Find(&books)
	} else {
		db.Order("cost asc").Find(&books)
	}
	c.JSON(http.StatusOK, books)
}
