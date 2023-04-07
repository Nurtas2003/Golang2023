package pkg

import (
	_ "github.com/lib/pq"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DB() *gorm.DB {
	const_str := "user=postgres password=t$hZw!Kz dbname=assign3 sslmode=disable"
	db, err := gorm.Open(mysql.Open(const_str), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
	return db
}
