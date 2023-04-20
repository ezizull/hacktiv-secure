package book

import (
	bookDomain "secure/challenge-3/domain/book"
	bookRepository "secure/challenge-3/infrastructure/repository/postgres/book"
)

type BookTesting interface {
	GetAll(page int64, limit int64) (*bookDomain.PaginationResultBook, error)
	UserGetAll(page int64, userId int, limit int64) (*bookDomain.PaginationResultBook, error)
	GetByID(id int) (*bookDomain.Book, error)
	UserGetByID(id int, userId int) (*bookDomain.Book, error)
	Create(book *NewBook) (*bookDomain.Book, error)
	GetByMap(medicineMap map[string]interface{}) (*bookDomain.Book, error)
	Delete(id int) (err error)
	Update(id uint, medicineMap map[string]interface{}) (*bookDomain.Book, error)
}

func NewTesting(bookTest bookRepository.BookTesting) BookTesting {
	return &Service{
		bookTesting: bookTest,
	}
}
