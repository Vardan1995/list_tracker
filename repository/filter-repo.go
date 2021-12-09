package repository

import (
	"github.com/Vardan1995/list_tracker/database"
	"github.com/Vardan1995/list_tracker/entity"
)

type filter_repo struct{}

func NewFilterRepository() FilterRepository {
	return &filter_repo{}
}

func (*filter_repo) SaveFilter(filter *entity.Filter) error {
	db := database.DB
	if err := db.Create(&filter).Error; err != nil {
		return err
	}
	return nil
}

func (*filter_repo) FindFilter(filter *entity.Filter, id uint) error {
	db := database.DB
	if err := db.First(&filter).Where("id = ?", id).Error; err != nil {
		return err
	}
	return nil
}

func (*filter_repo) FindFiltersByUserId(filters *[]entity.Filter, user_id uint) error {
	db := database.DB
	if err := db.Where("user_id = ?", user_id).Find(&filters).Error; err != nil {
		return err
	}
	return nil
}

func (*filter_repo) FindFilters(filters *[]entity.Filter) error {
	db := database.DB
	if err := db.Find(&filters).Select("link").Find(&filters).Error; err != nil {
		return err
	}
	return nil
}
func (*filter_repo) DeleteFilter(filter *entity.Filter, id uint) error {
	db := database.DB
	if err := db.Delete(&filter).Where("id = ?", id).Error; err != nil {
		return err
	}
	return nil
}
