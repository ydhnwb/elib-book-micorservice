package repository

import (
	"github.com/ydhnwb/elib-book-microservice/domain/entity"
	"gorm.io/gorm"
)

//BookRepository is a contract
type BookRepository interface {
	InsertBook(b entity.Book) entity.Book
	UpdateBook(b entity.Book) entity.Book
	DeleteBook(b entity.Book)
	AllBook() []entity.Book
	FindBookByID(bookID string) entity.Book
}

type bookConnection struct {
	connection *gorm.DB
}

//NewBookRepository creates an instance BookRepository
func NewBookRepository(dbConn *gorm.DB) BookRepository {
	return &bookConnection{
		connection: dbConn,
	}
}

func (db *bookConnection) InsertBook(b entity.Book) entity.Book {
	db.connection.Save(&b)
	return b
}

func (db *bookConnection) UpdateBook(b entity.Book) entity.Book {
	db.connection.Save(&b)
	return b
}

func (db *bookConnection) DeleteBook(b entity.Book) {
	db.connection.Delete(&b)
}

func (db *bookConnection) FindBookByID(bookID string) entity.Book {
	var book entity.Book
	db.connection.Find(&book, bookID)
	return book
}

func (db *bookConnection) AllBook() []entity.Book {
	var books []entity.Book
	db.connection.Find(&books)
	return books
}
