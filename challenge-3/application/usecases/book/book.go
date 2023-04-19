package book

import (
	bookDomain "secure/challenge-3/domain/book"
	bookRepository "secure/challenge-3/infrastructure/repository/postgres/book"
	userRepository "secure/challenge-3/infrastructure/repository/postgres/user"
)

// Service is a struct that contains the repository implementation for book use case
type Service struct {
	BookRepository bookRepository.Repository
	UserRepository userRepository.Repository
}

// GetAll is a function that returns all medicines
func (s *Service) GetAll(page int64, limit int64) (*PaginationResultBook, error) {

	all, err := s.BookRepository.GetAll(page, limit)
	if err != nil {
		return nil, err
	}

	return &PaginationResultBook{
		Data:       all.Data,
		Total:      all.Total,
		Limit:      all.Limit,
		Current:    all.Current,
		NextCursor: all.NextCursor,
		PrevCursor: all.PrevCursor,
		NumPages:   all.NumPages,
	}, nil
}

// UserGetAll is a function that returns all medicines
func (s *Service) UserGetAll(page int64, userId int, limit int64) (*PaginationResultBook, error) {

	all, err := s.BookRepository.UserGetAll(page, userId, limit)
	if err != nil {
		return nil, err
	}

	return &PaginationResultBook{
		Data:       all.Data,
		Total:      all.Total,
		Limit:      all.Limit,
		Current:    all.Current,
		NextCursor: all.NextCursor,
		PrevCursor: all.PrevCursor,
		NumPages:   all.NumPages,
	}, nil
}

// GetByID is a function that returns a book by id
func (s *Service) GetByID(id int) (*bookDomain.Book, error) {
	return s.BookRepository.GetByID(id)
}

// UserGetByID is a function that returns a book by id
func (s *Service) UserGetByID(id int, userId int) (*bookDomain.Book, error) {
	return s.BookRepository.UserGetByID(id, userId)
}

// Create is a function that creates a book
func (s *Service) Create(book *NewBook) (*bookDomain.Book, error) {

	bookModel := book.toDomainMapper()

	return s.BookRepository.Create(bookModel)
}

// GetByMap is a function that returns a book by map
func (s *Service) GetByMap(medicineMap map[string]interface{}) (*bookDomain.Book, error) {
	return s.BookRepository.GetOneByMap(medicineMap)
}

// Delete is a function that deletes a book by id
func (s *Service) Delete(id int) error {
	return s.BookRepository.Delete(id)
}

// Update is a function that updates a book by id
func (s *Service) Update(id uint, medicineMap map[string]interface{}) (*bookDomain.Book, error) {
	return s.BookRepository.Update(id, medicineMap)
}
