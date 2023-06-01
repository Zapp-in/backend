package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	AuthorID    string `gorm:"size:255;"`
	MusicUrl    string `gorm:"size:255;"`
	Name        string `gorm:"size:255;"`
	Description string `gorm:"size:255;"`
	Genre       string `gorm:"size:255;"`
	Label       string `gorm:"size:255;"`
}
