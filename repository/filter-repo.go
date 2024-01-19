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
	return db.Create(&filter).Error
}

func (*filter_repo) FindFilter(filter *entity.Filter, id uint) error {
	db := database.DB
	return db.First(&filter).Where("id = ?", id).Error
}

func (*filter_repo) FindFiltersByUserId(filters *[]entity.Filter, user_id uint) error {
	db := database.DB
	return db.Where("user_id = ?", user_id).Find(&filters).Error
}

func (*filter_repo) FindFilters(filters *[]entity.Filter) error {
	db := database.DB
	return db.Find(&filters).Select("link").Find(&filters).Error
}
func (*filter_repo) DeleteFilter(filter *entity.Filter, id uint) error {
	db := database.DB
	return db.Delete(&filter).Where("id = ?", id).Error
}
