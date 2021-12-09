package repository

import (
	"github.com/Vardan1995/filter_notifyer/database"
	"github.com/Vardan1995/filter_notifyer/entity"
)

type user_repo struct{}

func NewUserRepository() UserRepository {
	return &user_repo{}
}

func (*user_repo) SaveUser(user *entity.User) error {
	db := database.DB
	if err := db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func (*user_repo) FindUser(user *entity.User, id uint) error {
	db := database.DB
	if err := db.First(&user).Where("id = ?", id).Error; err != nil {
		return err
	}
	return nil
}
func (*user_repo) FindUserByUsernameAndEmail(user *entity.User, username, email string) error {
	db := database.DB
	if err := db.Where("username = ? and email = ?", username, email).First(&user).Error; err != nil {
		return err
	}
	return nil
}

func (*user_repo) FindUsers(users *[]entity.User) error {
	db := database.DB
	if err := db.Find(&users).Error; err != nil {
		return err
	}
	return nil
}

func (*user_repo) DeleteUser(user *entity.User, id uint) error {
	db := database.DB
	if err := db.Delete(&user).Where("id = ?", id).Error; err != nil {
		return err
	}
	return nil
}
