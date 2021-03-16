package interfaces

import "Users/core"

type IUsersRepository interface {
	CreateUser(user core.User) bool
	GetAllUsers() []*core.User
	GetUserById(user int) *core.User
	DeleteUser(user core.User) bool
	UpdateUser(user core.User) bool
}