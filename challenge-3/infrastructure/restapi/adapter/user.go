package adapter

import (
	userService "secure/challenge-3/application/usecases/user"
	roleRepository "secure/challenge-3/infrastructure/repository/postgres/role"
	userRepository "secure/challenge-3/infrastructure/repository/postgres/user"
	userController "secure/challenge-3/infrastructure/restapi/controllers/user"

	"gorm.io/gorm"
)

// UserAdapter is a function that returns a user controller
func UserAdapter(db *gorm.DB) *userController.Controller {
	uRepository := userRepository.Repository{DB: db}
	rRepository := roleRepository.Repository{DB: db}

	service := userService.Service{UserRepository: uRepository, RoleRepository: rRepository}
	return &userController.Controller{UserService: service}
}
