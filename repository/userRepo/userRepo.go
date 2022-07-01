package userrepo

import (
	"errors"
	"fmt"
	"test/delivery/view/request"
	"test/delivery/view/response"
	"test/entity"

	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type UserRepository struct {
	Db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		Db: db,
	}
}

func (u *UserRepository) InsertAdmin(newUser entity.Admin) (response.Admin, error) {

	err := u.Db.Create(&newUser).Error
	if err != nil {
		return response.Admin{}, err
	}
	var user response.Admin
	copier.Copy(&user, &newUser)
	return user, nil
}
func (u *UserRepository) InsertUser(newUser entity.User) (response.User, error) {

	err := u.Db.Create(&newUser).Error
	if err != nil {
		return response.User{}, err
	}
	var User response.User
	copier.Copy(&User, &newUser)
	return User, nil
}

func (u *UserRepository) LoginAdmin(login request.Login) (entity.Admin, error) {
	var user entity.Admin
	fmt.Printf(login.Email)
	err := u.Db.Where("email = ?", login.Email).First(&user).Error
	if err != nil {
		return entity.Admin{}, errors.New("Email yang anda masukan salah")
	}

	return user, nil
}
func (u *UserRepository) Login(login request.Login) (entity.User, error) {
	var user entity.User
	err := u.Db.Where("email=?", login.Email).First(&user).Error
	if err != nil {
		return entity.User{}, errors.New("Email yang anda masukan salah")
	}
	return user, nil
}

func (u *UserRepository) GetUser(userID uint) (response.User, error) {
	var user entity.User
	err := u.Db.Where("id = ?", userID).First(&user).Error
	if err != nil {
		return response.User{}, err
	}
	var User response.User
	copier.Copy(&User, &user)
	return User, nil
}
func (u *UserRepository) GetAdmin(userID uint) (response.Admin, error) {
	var user entity.Admin
	err := u.Db.Where("id = ?", userID).First(&user).Error
	if err != nil {
		return response.Admin{}, err
	}
	var User response.Admin
	copier.Copy(&User, &user)
	return User, nil
}

func (u *UserRepository) GetAllUser(offsite int) ([]response.User, error) {
	var users []entity.User
	err := u.Db.Limit(4).Offset(offsite).Find(&users).Error
	if err != nil {
		return []response.User{}, err
	}
	var Users []response.User
	copier.Copy(&Users, &users)
	return Users, nil
}
func (u *UserRepository) Search(title string, offsite int) ([]response.User, error) {
	var users []entity.User
	err := u.Db.Where("nama LIKE ?", "%"+title+"%").Limit(4).Offset(offsite).Find(&users).Error
	if err != nil {
		return []response.User{}, err
	}
	var Users []response.User
	copier.Copy(&Users, &users)
	return Users, nil
}

func (u *UserRepository) UpdateUser(userID uint, Newuser request.UpdateUser) (response.User, error) {
	var user entity.User
	err := u.Db.Where("id=?", userID).First(&user).Updates(&Newuser).Error
	if err != nil {
		return response.User{}, err
	}
	var Users response.User
	copier.Copy(&Users, &user)
	return Users, nil
}

func (u *UserRepository) DeleteUser(userID uint) error {
	var user entity.User
	err := u.Db.Where("id=?", userID).Delete(&user).Error
	if err != nil {
		return err
	}
	return nil
}
