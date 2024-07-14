package repository

import (
	"gorm.io/gorm"

	m "GoDeveloperTest/models"
)

type bookRepositoryDB struct {
	db *gorm.DB
}

func NewBookDB(db *gorm.DB) *bookRepositoryDB {
	return &bookRepositoryDB{db: db}
}


func (r *bookRepositoryDB) GetAll() ([]m.Book, error) {
	var books []m.Book

	result := r.db.Find(&books)
	if result.Error != nil {
		return nil, result.Error
	}
	return books, nil
}

func (r *bookRepositoryDB) GetById(id int) (*m.Book, error)  {
	var book m.Book

	result := r.db.First(&book, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &book, nil
}

func (r *bookRepositoryDB) CreateOne(book *m.Book) error {
	result := r.db.Create(book)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *bookRepositoryDB) UpdateOne(book m.Book) error  {
	result := r.db.Save(&book)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *bookRepositoryDB) DeleteOne(id int) error  {
	var book m.Book

	result := r.db.Delete(&book, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}