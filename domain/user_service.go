package domain

import (
	"context"

	"github.com/openmesh/kitgen/models"
)

type UserService interface {
	CreateUser(ctx context.Context, req CreateUserRequest) CreateUserResponse
	CreateUsersWithArrayInput(ctx context.Context, req CreateUsersWithArrayInputRequest) CreateUsersWithArrayInputResponse
	CreateUsersWithListInput(ctx context.Context, req CreateUsersWithListInputRequest) CreateUsersWithListInputResponse
	DeleteUser(ctx context.Context, req DeleteUserRequest) DeleteUserResponse
	GetUserByName(ctx context.Context, req GetUserByNameRequest) GetUserByNameResponse
	LoginUser(ctx context.Context, req LoginUserRequest) LoginUserResponse
	LogoutUser(ctx context.Context, req LogoutUserRequest) LogoutUserResponse
	UpdateUser(ctx context.Context, req UpdateUserRequest) UpdateUserResponse
}

type UserService interface {
	CreateUser(ctx context.Context, body models.User)
	CreateUsersWithArrayInput(ctx context.Context, body []*models.User)
	CreateUsersWithListInput(ctx context.Context, body []*models.User)
	DeleteUser(ctx context.Context, username string)
	GetUserByName(ctx context.Context, username string) models.User
	LoginUser(ctx context.Context, password string, username string) string
	LogoutUser(ctx context.Context)
	UpdateUser(ctx context.Context, body models.User, username string)
}
