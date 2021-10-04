package domain

import (
	"context"
)

type UserService interface {
	CreateUser(ctx context.Context, body User) error
	CreateUsersWithArrayInput(ctx context.Context, body User) error
	CreateUsersWithListInput(ctx context.Context, body User) error
	DeleteUser(ctx context.Context, username string) error
	GetUserByName(ctx context.Context, username string) (User, error)
	LoginUser(ctx context.Context, password string, username string) (string, error)
	LogoutUser(ctx context.Context) error
	UpdateUser(ctx context.Context, body User, username string) error
}
