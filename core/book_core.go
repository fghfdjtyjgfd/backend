package core

import (
	"log"

	m "GoDeveloperTest/models"
	repo "GoDeveloperTest/repository"
)

type bookService struct {
	bookRepo repo.BookRepository
}

func NewBookService(bookRepo repo.BookRepository) *bookService {
	return &bookService{bookRepo: bookRepo}
}

func (s *bookService) GetBooks() ([]m.Book, error) {
	books, err := s.bookRepo.GetAll()
	if err != nil {
		log.Print(err)
	}
	return books, nil
}

func (s *bookService) GetBookById(id int) (*m.Book, error) {
	book, err := s.bookRepo.GetById(id)
	if err != nil {
		log.Print(err)
	}
	return book, nil
}

func (s *bookService) CreateBook(book m.Book) error {
	err := s.bookRepo.CreateOne(&book)
	if err != nil {
		log.Print(err)
	}
	return nil
}

func (s *bookService) UpdateBook(book m.Book) error {
	err := s.bookRepo.UpdateOne(book)
	if err != nil {
		log.Print(err)
	}
	return nil
}

func (s *bookService) DeleteBook(id int) error {
	err := s.bookRepo.DeleteOne(id)
	if err != nil {
		log.Print(err)
	}
	return nil
}
