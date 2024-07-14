package core

import (
	m "GoDeveloperTest/models"
)

type BookService interface {
	GetBooks () ([]m.Book, error)
	GetBookById (id int) (*m.Book, error)
	CreateBook (book m.Book) error
	UpdateBook (book m.Book) error
	DeleteBook (id int) error
}