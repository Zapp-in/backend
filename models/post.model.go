package models

import "gorm.io/gorm"



type Post struct {
	gorm.Model
	AuthorID    string `gorm:"size:255;not null;unique;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	MusicUrl    string `gorm:"size:255;"`
	Name        string `gorm:"size:255;"`
	Description string `gorm:"size:255;"`
	Genre       string `gorm:"size:255;"`
	Label       string `gorm:"size:255;"`
}
