// Package user contains the user controller
package user

import (
	userUseCase "secure/challenge-2/application/usecases/user"
	domainRole "secure/challenge-2/domain/role"
	userDomain "secure/challenge-2/domain/user"
)

func userRoleDomainToResponseMapper(userRoleDomain *domainRole.User) (createUserRoleResponse *ResponseUserRole) {
	createUserRoleResponse = &ResponseUserRole{ID: userRoleDomain.ID, UserName: userRoleDomain.UserName,
		Email: userRoleDomain.Email, FirstName: userRoleDomain.FirstName, LastName: userRoleDomain.LastName,
		Status: userRoleDomain.Status, Role: userRoleDomain.Role}

	return
}

func domainToResponseMapper(userDomain *userDomain.User) (createUserResponse *ResponseUser) {
	createUserResponse = &ResponseUser{ID: userDomain.ID, UserName: userDomain.UserName,
		Email: userDomain.Email, FirstName: userDomain.FirstName, LastName: userDomain.LastName,
		Status: userDomain.Status, CreatedAt: userDomain.CreatedAt, UpdatedAt: userDomain.UpdatedAt}

	return
}

func arrayDomainToResponseMapper(usersDomain *[]userDomain.User) *[]ResponseUser {
	usersResponse := make([]ResponseUser, len(*usersDomain))
	for i, user := range *usersDomain {
		usersResponse[i] = *domainToResponseMapper(&user)
	}
	return &usersResponse
}

func toUsecaseMapper(user *NewUserRequest) *userUseCase.NewUser {
	return &userUseCase.NewUser{UserName: user.UserName, Password: user.Password, Email: user.Email, FirstName: user.FirstName, LastName: user.LastName, RoleID: user.RoleID}

}
