package book

import (
	domainBook "secure/challenge-3/domain/book"
	bookRepository "secure/challenge-3/infrastructure/repository/postgres/book"
)

type BookMocks interface {
	GetAll(page int64, limit int64) (*PaginationResultBook, error)
	UserGetAll(page int64, userId int, limit int64) (*PaginationResultBook, error)
	Create(newBook *domainBook.Book) (createdBook *domainBook.Book, err error)
	GetByID(id int) (*domainBook.Book, error)
	UserGetByID(id int, userId int) (*domainBook.Book, error)
	GetOneByMap(bookMap map[string]interface{}) (*domainBook.Book, error)
	Update(id uint, bookMap map[string]interface{}) (*domainBook.Book, error)
	Delete(id int) (err error)
}

type bookTesting struct {
	bookTest bookRepository.BookMocks
}

func NewTesting(bookTest bookRepository.BookMocks) BookMocks {
	return nil
}
