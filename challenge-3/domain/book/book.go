package book

import (
	"time"
)

type Book struct {
	ID          uint   `json:"id" gorm:"primary_key"`
	Title       string `json:"title" example:"book title"`
	UserID      int    `json:"user_id" gorm:"index"`
	Description string `json:"description" example:"book description"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
}

func (*Book) TableName() string {
	return "books"
}
