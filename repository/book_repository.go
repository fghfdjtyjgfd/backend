package repository

import (
	m "GoDeveloperTest/models"
)

type BookRepository interface {
	GetAll() ([]m.Book, error)
	GetById (id int) (*m.Book, error)
	CreateOne (book *m.Book) error
	UpdateOne (book m.Book) error
	DeleteOne (id int) error
	
}