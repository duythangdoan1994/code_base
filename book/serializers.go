package book

import (
	"github.com/jinzhu/gorm"
)

func (b *Book) TableName() string {
	return "books"
}

func InsertBook(db *gorm.DB, b *Book) (err error) {
	if err = db.Save(b).Error; err != nil {
		return err
	}
	return nil
}

func GetAllBook(db *gorm.DB, b *[]Book) (err error) {
	if err = db.Order("id desc").Find(b).Error; err != nil {
		return err
	}
	return nil
}

func OneBookGetting(db *gorm.DB, ids int, b *Book) (err error) {
	if err := db.Where("codes = ?", ids).First(&b).Error; err != nil {
		return err
	}
	return nil
}

func UpdateBook(db *gorm.DB, b *Book) (err error) {
	if err = db.Save(b).Error; err != nil {
		return err
	}
	return nil
}

func DeletedBook(db *gorm.DB, b *Book) (err error) {
	if err = db.Delete(b).Error; err != nil {
		return err
	}
	return nil
}