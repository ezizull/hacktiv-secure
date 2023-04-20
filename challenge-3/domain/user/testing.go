package user

/* Testing */
type Testing interface {
	Get(int) (*User, error)
	GetAll() ([]*User, error)
	Create(*User) error
	GetByMap(map[string]interface{}) map[string]interface{}
	GetByID(int) (*User, error)
	Delete(int) error
	Update(int, map[string]interface{}) (*User, error)
}
