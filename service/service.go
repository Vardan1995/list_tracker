package service

import "github.com/Vardan1995/list_tracker/entity"

type UserService interface {
	CreateUser(user *entity.User) (*entity.User, string, error)
	GetUser(user *entity.User, id uint) (*entity.User, error)
	GetUsers(users *[]entity.User) error

	LoginUser(user *entity.User, username, email string) (string, error)
	// GetUsers(user *[]entity.User) (*[]entity.User, error)
	// DeleteUser(user *entity.User) error
}

type ArchiveService interface {
	CreateIfNotExists(email, product_id string) bool
	// CreateArchive(archive *entity.Archive) (*entity.Archive, error)
	// GetArchive(archive *entity.Archive) (*entity.Archive, error)
	// GetArchives(archive *[]entity.Archive) (*[]entity.Archive, error)
	// DeleteArchive(archive *entity.Archive) error
}

type FilterService interface {
	CreateFilter(filter *entity.Filter) (*entity.Filter, error)
	GetUserFilter(filter *[]entity.Filter, id uint) error
	GetFilters(filter *[]entity.Filter) error
	// DeleteFilter(filter *entity.Filter) error
}
