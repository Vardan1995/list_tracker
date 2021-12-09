package repository

import "github.com/Vardan1995/filter_notifyer/entity"

type UserRepository interface {
	SaveUser(user *entity.User) error
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
type FilterRepository interface {
	SaveFilter(filter *entity.Filter) error
	FindFilter(filter *entity.Filter, id uint) error
	FindFiltersByUserId(filters *[]entity.Filter, user_id uint) error
	FindFilters(filters *[]entity.Filter) error
	DeleteFilter(filter *entity.Filter, id uint) error
}
