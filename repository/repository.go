package repository

import "github.com/Vardan1995/list_tracker/entity"

type UserRepository interface {
	SaveUser(user *entity.User) error
	UserWithPreferences(user *entity.User) error
	FindUserByUsernameAndEmail(user *entity.User, username, email string) error
	FindUser(user *entity.User, id uint) error
	FindUsers(user *[]entity.User) error
	DeleteUser(user *entity.User, id uint) error
}
type ArchiveRepository interface {
	SaveArchive(archive *entity.Archive) (*entity.Archive, error)
	FindProductFromArchive(archive *entity.Archive, email, product_id string) error
	DeleteArchive(archive *entity.Archive, id uint) error
}
type PreferenceRepository interface {
	SavePreference(preference *entity.Preference) error
	FindPreference(preference *entity.Preference, id uint) error
	FindPreferencesByUserId(preferences *[]entity.Preference, user_id uint) error
	FindPreferences(preferences *[]entity.Preference) error
	DeletePreference(preference *entity.Preference, id uint) error
}
