// Package user contains the business logic for the user entity
package user

import (
	domainRole "secure/challenge-2/domain/role"
	domainUser "secure/challenge-2/domain/user"
)

// toDomainMapper function to convert user repo to user domain
func (user *User) toDomainMapper() *domainUser.User {
	return &domainUser.User{
		ID:           user.ID,
		UserName:     user.UserName,
		Email:        user.Email,
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		Status:       user.Status,
		HashPassword: user.HashPassword,
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
	}
}

// fromDomainMapper function to convert user domain to user repo
func fromDomainMapper(user *domainUser.User) *User {
	return &User{
		ID:           user.ID,
		UserName:     user.UserName,
		Email:        user.Email,
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		Status:       user.Status,
		HashPassword: user.HashPassword,
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
		RoleID:       user.RoleID,
	}
}

// arrayToDomainMapper function to convert list user domain to list user repo
func arrayToDomainMapper(users *[]User) *[]domainUser.User {
	usersDomain := make([]domainUser.User, len(*users))
	for i, user := range *users {
		usersDomain[i] = *user.toDomainMapper()
	}

	return &usersDomain
}

// toDomainMapper function to convert role of user role repo to role domain
func (userRole *UserRole) toRoleDomainMapper() *domainRole.Role {
	return &domainRole.Role{
		ID:        userRole.ID,
		Name:      userRole.Role.Name,
		CreatedAt: userRole.CreatedAt,
		UpdatedAt: userRole.UpdatedAt,
	}
}

// toDomainMapper function to convert userRole repo to userRole domain
func (userRole *UserRole) toDomainMapper() *domainRole.User {
	return &domainRole.User{
		ID:        userRole.ID,
		UserName:  userRole.UserName,
		Email:     userRole.Email,
		FirstName: userRole.FirstName,
		LastName:  userRole.LastName,
		Status:    userRole.Status,
		RoleID:    userRole.RoleID,
		Role:      *userRole.toRoleDomainMapper(),
	}
}
