package types

import (
	"gorm.io/gorm"
	"log"
	"time"
)

type Article struct {
	ID          uint      `gorm:"primary" json:"id"`
	Title       string    `json:"title"`
	Content     string    `json:"content"`
	Category    string    `json:"category"`
	Author      string    `json:"author"`
	PublishedAt time.Time `json:"published_at"`
}
type Category struct {
	ID   uint   `gorm:"primary" json:"id"`
	Name string `json:"name"`
}

func Migration(db *gorm.DB) error {
	err := db.AutoMigrate(&Article{})
	if err != nil {
		log.Fatal("Cannot migrate Article")
	}
	err = db.AutoMigrate(&Category{})
	if err != nil {
		log.Fatal("Cannot migrate Category")
	}
	return nil
}
