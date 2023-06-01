package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email  string `gorm:"size:255;not null;unique" json:"email"`
	Name   string `gorm:"size:255;" json:"name"`
	Musics []Post `gorm:"foreignKey:AuthorID"`
}
