package book

import "time"

// TestBook is a struct that contains the testing book model
type TestBook struct {
	ID          uint   `json:"id" gorm:"primary_key"`
	Title       string `json:"title" example:"book title"`
	UserID      int    `json:"user_id" gorm:"index"`
	Description string `json:"description" example:"book description"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
}

// TableName overrides the table name used by TestBook to `test_books`
func (*TestBook) TableName() string {
	return "test_books"
}

/* Testing */
type Testing interface {
	Get(int) (*Book, error)
	GetAll() ([]*Book, error)
	Create(*Book) error
	GetByMap(map[string]interface{}) map[string]interface{}
	GetByID(int) (*Book, error)
	Delete(int) error
	Update(int, map[string]interface{}) (*Book, error)
}
