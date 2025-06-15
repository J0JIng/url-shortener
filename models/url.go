package models

import(
	"gorm.io/gorm"
)

type URL struct {
	gorm.Model     
	OriginalURL string     `gorm:"column:original_url"`
	ShortenURL  string     `gorm:"column:shorten_url"`      
}