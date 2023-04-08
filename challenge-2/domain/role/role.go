package role

import "time"

// User is a struct that contains the user information
type User struct {
	ID        int
	UserName  string
	Email     string
	FirstName string
	LastName  string
	Status    bool
	RoleID    string
	Role      Role `gorm:"foreignKey:ID;references:RoleID"`
}

// Role is a struct that contains the role information
type Role struct {
	ID        int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
