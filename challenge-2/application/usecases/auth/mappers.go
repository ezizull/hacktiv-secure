// Package auth provides the use case for authentication
package auth

import (
	roleDomain "secure/challenge-2/domain/role"
	userDomain "secure/challenge-2/domain/user"
)

func secAuthUserMapper(domainUser *userDomain.User, authInfo *Auth) *SecurityAuthenticatedUser {
	return &SecurityAuthenticatedUser{
		Data: DataUserAuthenticated{
			UserName:  domainUser.UserName,
			Email:     domainUser.Email,
			FirstName: domainUser.FirstName,
			LastName:  domainUser.LastName,
			ID:        domainUser.ID,
			RoleID:    domainUser.RoleID,
		},
		Security: DataSecurityAuthenticated{
			JWTAccessToken:            authInfo.AccessToken,
			JWTRefreshToken:           authInfo.RefreshToken,
			ExpirationAccessDateTime:  authInfo.ExpirationAccessDateTime,
			ExpirationRefreshDateTime: authInfo.ExpirationRefreshDateTime,
		},
	}

}

func secAuthUserRoleMapper(domainUserRole *roleDomain.User, authInfo *Auth) *SecurityAuthenticatedUser {
	return &SecurityAuthenticatedUser{
		Data: DataUserAuthenticated{
			ID:        domainUserRole.ID,
			UserName:  domainUserRole.UserName,
			Email:     domainUserRole.Email,
			FirstName: domainUserRole.FirstName,
			LastName:  domainUserRole.LastName,
			RoleID:    domainUserRole.RoleID,
			Role:      domainUserRole.Role,
		},
		Security: DataSecurityAuthenticated{
			JWTAccessToken:            authInfo.AccessToken,
			JWTRefreshToken:           authInfo.RefreshToken,
			ExpirationAccessDateTime:  authInfo.ExpirationAccessDateTime,
			ExpirationRefreshDateTime: authInfo.ExpirationRefreshDateTime,
		},
	}

}
