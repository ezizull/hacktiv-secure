// Package user provides the use case for user
package user

import (
	roleDomain "secure/challenge-3/domain/role"
	userDomain "secure/challenge-3/domain/user"
	roleRepository "secure/challenge-3/infrastructure/repository/postgres/role"
	userRepository "secure/challenge-3/infrastructure/repository/postgres/user"

	"golang.org/x/crypto/bcrypt"
)

// Service is a struct that contains the repository implementation for user use case
type Service struct {
	UserRepository userRepository.Repository
	RoleRepository roleRepository.Repository
}

// GetAll is a function that returns all users
func (s *Service) GetAll() (*[]userDomain.User, error) {
	return s.UserRepository.GetAll()
}

// GetWithRole is a function that returns a user with role by id
func (s *Service) GetWithRole(id int) (*roleDomain.User, error) {
	return s.UserRepository.GetWithRole(id)
}

// GetByID is a function that returns a user by id
func (s *Service) GetByID(id int) (*userDomain.User, error) {
	return s.UserRepository.GetByID(id)
}

// Create is a function that creates a new user
func (s *Service) Create(newUser NewUser) (*userDomain.User, error) {

	_, err := s.RoleRepository.GetByID(newUser.RoleID)
	if err != nil {
		return &userDomain.User{}, err
	}

	domain := newUser.toDomainMapper()

	hash, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		return &userDomain.User{}, err
	}
	domain.HashPassword = string(hash)

	return s.UserRepository.Create(domain)
}

// GetOneByMap is a function that returns a user by map
func (s *Service) GetOneByMap(userMap map[string]interface{}) (*userDomain.User, error) {
	return s.UserRepository.GetOneByMap(userMap)
}

// Delete is a function that deletes a user by id
func (s *Service) Delete(id int) error {
	return s.UserRepository.Delete(id)
}

// Update is a function that updates a user by id
func (s *Service) Update(id int, userMap map[string]interface{}) (*userDomain.User, error) {
	return s.UserRepository.Update(id, userMap)
}
