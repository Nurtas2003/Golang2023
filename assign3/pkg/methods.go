package pkg

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"gorm.io/gorm"
	"net/http"
	"strings"
)

type Connection struct {
	DB *gorm.DB
}

func (c *Connection) GetAllBooks(w http.ResponseWriter, r *http.Request) {
	books := make([]Book, 0)
	c.DB.Find(&books)
	err := json.NewEncoder(w).Encode(books)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (c *Connection) GetBookByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var book Book
	if err := c.DB.First(&book, params["id"]).Error; err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(book)
}

func (c *Connection) AddBook(w http.ResponseWriter, r *http.Request) {
	var book []Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	c.DB.Create(&book)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(book)
}

func (c *Connection) UpdateBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var book Book
	if err := c.DB.Where("id = ?", params["id"]).First(&book).Error; err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := c.DB.Save(&book).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(book)
}

func (c *Connection) DeleteBookByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	c.DB.Delete(&Book{}, params["id"])
}

func (c *Connection) SearchBookByTitle(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var books []Book
	c.DB.Where("title LIKE ?", "%"+params["title"]+"%").Find(&books)
	json.NewEncoder(w).Encode(books)
}

func (c *Connection) GetSortedBooks(w http.ResponseWriter, r *http.Request) {
	var book []Book
	sort := r.URL.Query().Get("sort")
	parts := strings.Split(sort, "-")
	sorting := strings.Join(parts, " ")
	fmt.Println(parts)
	if sorting == "" {
		sorting = "id asc"
	}
	if err := c.DB.Order(sorting).Find(&book).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(book); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (c *Connection) DescGetSortedBooks(w http.ResponseWriter, r *http.Request) {
	var book []Book
	sort := r.URL.Query().Get("sort")
	parts := strings.Split(sort, "-")
	sorting := strings.Join(parts, " ")
	fmt.Println(parts)
	if sorting == "" {
		sorting = "id desc"
	}
	if err := c.DB.Order(sorting).Find(&book).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(book); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
