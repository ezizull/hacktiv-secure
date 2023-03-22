package book

import (
	bookDomain "microservices/challenge-4-advance/domain/book"
	bookRepository "microservices/challenge-4-advance/infrastructure/repository/book"
)

// Service is a struct that contains the repository implementation for medicine use case
type Service struct {
	BookRepository bookRepository.Repository
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

// GetByID is a function that returns a medicine by id
func (s *Service) GetByID(id int) (*bookDomain.Book, error) {
	return s.BookRepository.GetByID(id)
}

// Create is a function that creates a medicine
func (s *Service) Create(medicine *NewBook) (*bookDomain.Book, error) {
	medicineModel := medicine.toDomainMapper()
	return s.BookRepository.Create(medicineModel)
}

// GetByMap is a function that returns a medicine by map
func (s *Service) GetByMap(medicineMap map[string]interface{}) (*bookDomain.Book, error) {
	return s.BookRepository.GetOneByMap(medicineMap)
}

// Delete is a function that deletes a medicine by id
func (s *Service) Delete(id int) error {
	return s.BookRepository.Delete(id)
}

// Update is a function that updates a medicine by id
func (s *Service) Update(id uint, medicineMap map[string]interface{}) (*bookDomain.Book, error) {
	return s.BookRepository.Update(id, medicineMap)
}
