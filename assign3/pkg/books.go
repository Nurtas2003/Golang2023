package pkg

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	ID          uint    `gorm:"primary key"`
	Title       string  `gorm:"type:varchar(255);not null"`
	Description string  `gorm:"type:text"`
	Cost        float32 `gorm:"not null"`
}
