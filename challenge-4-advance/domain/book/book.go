package book

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	ID          uint      `json:"id" gorm:"primary_key"`
	Title       string    `json:"title" example:"book title"`
	Author      string    `json:"author" example:"mr. author"`
	Description string    `json:"description" example:"book description"`
	CreatedAt   time.Time `json:"created_at,omitempty" gorm:"-"`
	UpdatedAt   time.Time `json:"updated_at,omitempty" example:"2021-02-24 20:19:39"`
	DeletedAt   *gorm.DeletedAt
}

// Service is a interface that contains the methods for the book service
type Service interface {
	Get(int) (*Book, error)
	GetAll() ([]*Book, error)
	Create(*Book) error
	GetByMap(map[string]interface{}) map[string]interface{}
	GetByID(int) (*Book, error)
	Delete(int) error
	Update(int, map[string]interface{}) (*Book, error)
}
