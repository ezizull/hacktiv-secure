package book

import (
	"time"
)

type TestBook struct {
	ID          uint   `json:"id" gorm:"primary_key"`
	Title       string `json:"title" example:"book title"`
	UserID      int    `json:"user_id" gorm:"index"`
	Description string `json:"description" example:"book description"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
}

func (*TestBook) TableName() string {
	return "test_books"
}

// ServiceTest is a interface that contains the methods for the book
type ServiceTest interface {
	Get(int) (*TestBook, error)
	GetAll() ([]*TestBook, error)
	Create(*TestBook) error
	GetByMap(map[string]interface{}) map[string]interface{}
	GetByID(int) (*TestBook, error)
	Delete(int) error
	Update(int, map[string]interface{}) (*TestBook, error)
}
