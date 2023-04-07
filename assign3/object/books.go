package object

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Cost        float32 `json:"cost"`
}
