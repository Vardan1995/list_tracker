package repository

import (
	"github.com/Vardan1995/list_tracker/database"
	"github.com/Vardan1995/list_tracker/entity"
)

type user_repo struct{}

var db = database.DB

func NewUserRepository() UserRepository {
	return &user_repo{}
}

func (*user_repo) SaveUser(user *entity.User) error {
	db := database.DB
	return db.Create(&user).Error
}

func (*user_repo) FindUser(user *entity.User, id uint) error {
	db := database.DB
	return db.First(&user).Where("id = ?", id).Error
}

func (*user_repo) UserWithPreferences(user *entity.User) error {
	db := database.DB
	return db.Model(&user).Preload("Preferences").Find(&user).Error

}

func (*user_repo) FindUserByUsernameAndEmail(user *entity.User, username, email string) error {
	db := database.DB
	return db.Where("username = ? and email = ?", username, email).First(&user).Error
}

func (*user_repo) FindUsers(users *[]entity.User) error {
	db := database.DB
	return db.Find(&users).Error
}

func (*user_repo) DeleteUser(user *entity.User, id uint) error {
	db := database.DB
	return db.Delete(&user).Where("id = ?", id).Error
}
