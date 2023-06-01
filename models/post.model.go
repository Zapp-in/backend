package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	AuthorID    string `gorm:"size:255;" json:"author_id"`
	Url         string `gorm:"size:255;" json:"url"`
	Name        string `gorm:"size:255;" json:"name"`
	Description string `gorm:"size:255;" json:"desc"`
	Genre       string `gorm:"size:255;" json:"genre"`
	Label       string `gorm:"size:255;" json:"label"`
}
