package models

import (
	"github.com/LuluBeatson/go-server/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name      string `json:"name"`
	Author    string `json:"author"`
	Publisher string `json:"publisher"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var book Book
	db := db.Where("ID = ?", Id).Find(&book)
	return &book, db
}

func DeleteBook(Id int64) (*Book, *gorm.DB) {
	var book Book
	db := db.Where("ID = ?", Id).Delete(&book)
	return &book, db
}
