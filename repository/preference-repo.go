package repository

import (
	"github.com/Vardan1995/list_tracker/database"
	"github.com/Vardan1995/list_tracker/entity"
)

type preference_repo struct{}

func NewPreferenceRepository() PreferenceRepository {
	return &preference_repo{}
}

func (*preference_repo) SavePreference(preference *entity.Preference) error {
	db := database.DB
	return db.Create(&preference).Error
}

func (*preference_repo) FindPreference(preference *entity.Preference, id uint) error {
	db := database.DB
	return db.First(&preference).Where("id = ?", id).Error
}

func (*preference_repo) FindPreferencesByUserId(preferences *[]entity.Preference, user_id uint) error {
	db := database.DB
	return db.Where("user_id = ?", user_id).Find(&preferences).Error
}

func (*preference_repo) FindPreferences(preferences *[]entity.Preference) error {
	db := database.DB
	return db.Find(&preferences).Select("link").Find(&preferences).Error
}
func (*preference_repo) DeletePreference(preference *entity.Preference, id uint) error {
	db := database.DB
	return db.Delete(&preference).Where("id = ?", id).Error
}
