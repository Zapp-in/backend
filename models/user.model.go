package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID     string `gorm:"primaryKey;size:255;"`
	Email  string `gorm:"size:255;not null;unique"`
	Name   string `gorm:"size:255;"`
	Musics []Post `gorm:"foreignKey:AuthorID"`
}
