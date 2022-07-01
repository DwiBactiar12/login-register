package userrepo

import (
	"test/delivery/view/request"
	"test/delivery/view/response"
	"test/entity"
)

type UserRepo interface {
	InsertUser(newUser entity.User) (response.User, error)
	GetUser(userID uint) (response.User, error)
	GetAdmin(userID uint) (response.Admin, error)
	GetAllUser(offsite int) ([]response.User, error)
	InsertAdmin(newUser entity.Admin) (response.Admin, error)
	Login(login request.Login) (entity.User, error)
	LoginAdmin(login request.Login) (entity.Admin, error)
	UpdateUser(userID uint, Newuser request.UpdateUser) (response.User, error)
	DeleteUser(userID uint) error
	Search(title string,offsite int) ([]response.User, error)
}
