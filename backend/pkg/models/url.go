package models

import (
	"log"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type URL struct {
	ID          uint `gorm:"primary_key"`
	OriginalURL string
	ShortURL    string `gorm:"unique"`
	CreatedAt   time.Time
	ExpiresAt   *time.Time
}

var DB *gorm.DB

func init() {
	var err error
	DB, err = gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=postgres password=veds2003 sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	DB.AutoMigrate(&URL{})
}
