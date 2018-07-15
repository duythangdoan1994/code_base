package book

import "github.com/jinzhu/gorm"

type Book struct {
	gorm.Model
	Codes int `gorm:"column:codes"`
	Name string `gorm:"column:name" json:"name"`
	Author string   `gorm:"column:author" json:"author"`
	Category string `gorm:"column:category" json:"category"`
}

func (b *Book) TableName() string {
	return "books"
}