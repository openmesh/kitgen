package endpoint

import (
	"github.com/go-kit/kit/endpoint"
	"github.com/openmesh/kitgen/domain"
)

type UserEndpoints struct {
	CreateUserEndpoint                endpoint.Endpoint
	CreateUsersWithArrayInputEndpoint endpoint.Endpoint
	CreateUsersWithListInputEndpoint  endpoint.Endpoint
	DeleteUserEndpoint                endpoint.Endpoint
	GetUserByNameEndpoint             endpoint.Endpoint
	LoginUserEndpoint                 endpoint.Endpoint
	LogoutUserEndpoint                endpoint.Endpoint
	UpdateUserEndpoint                endpoint.Endpoint
}

func MakeUserEndpoints(s domain.UserService) UserEndpoints {
	return UserEndpoints{
		CreateUserEndpoint:                MakeCreateUserEndpoint(s),
		CreateUsersWithArrayInputEndpoint: MakeCreateUsersWithArrayInputEndpoint(s),
		CreateUsersWithListInputEndpoint:  MakeCreateUsersWithListInputEndpoint(s),
		DeleteUserEndpoint:                MakeDeleteUserEndpoint(s),
		GetUserByNameEndpoint:             MakeGetUserByNameEndpoint(s),
		LoginUserEndpoint:                 MakeLoginUserEndpoint(s),
		LogoutUserEndpoint:                MakeLogoutUserEndpoint(s),
		UpdateUserEndpoint:                MakeUpdateUserEndpoint(s),
	}
}

type createUserRequest struct {
	Body
}

type createUsersWithArrayInputRequest struct {
	Body
}

type createUsersWithListInputRequest struct {
	Body
}

type deleteUserRequest struct {
	Username
}

type getUserByNameRequest struct {
	Username
}

type loginUserRequest struct {
	Password

	Username
}

type logoutUserRequest struct {
}

type updateUserRequest struct {
	Body

	Username
}
