package service

import (
	"github.com/Vardan1995/list_tracker/entity"
	"github.com/Vardan1995/list_tracker/repository"
)

var (
	preferenceRepository = repository.NewPreferenceRepository()
)

type preference_service struct{}

func NewPreferenceService() PreferenceService {
	return &preference_service{}
}

func (*preference_service) CreatePreference(preference *entity.Preference) (*entity.Preference, error) {
	if err := preferenceRepository.SavePreference(preference); err != nil {
		return nil, err
	}
	return preference, nil

}

func (*preference_service) GetUserPreference(preferences *[]entity.Preference, id uint) error {
	if err := preferenceRepository.FindPreferencesByUserId(preferences, id); err != nil {
		return err
	}
	return nil
}
func (*preference_service) GetPreferences(preferences *[]entity.Preference) error {
	if err := preferenceRepository.FindPreferences(preferences); err != nil {
		return err
	}
	return nil
}
