package book

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
