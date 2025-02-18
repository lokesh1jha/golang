package models

import (
	"github.com/jinzhu/gorm"
	"github.com/lokesh1jha/bookstore/pkg/config"
)

var db *gorm.DB

type Book struct {
	gorm.Model         // Embeds fields ID, CreatedAt, UpdatedAt, DeletedAt
	Title       string `json:"title"`       // Title of the book
	Author      string `json:"author"`      // Author of the book
	Publication string `json:"publication"` // Publication of the book
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
	var books []Book
	db.Find(&books)
	return books
}

func GetBookById(bookId int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID = ?", bookId).Find(&getBook)
	return &getBook, db
}

func DeleteBook(bookId int64) Book {
	var book Book
	db.Where("ID = ?", bookId).Delete(book)
	return book
}

func (b *Book) UpdateBook(bookId int64) *Book {
	db.Save(&b)
	return b
}
