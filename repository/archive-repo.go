package repository

import (
	"github.com/Vardan1995/filter_notifyer/database"
	"github.com/Vardan1995/filter_notifyer/entity"
)

type archive_repo struct{}

func NewArchiveRepository() ArchiveRepository {
	return &archive_repo{}
}

func (*archive_repo) SaveArchive(archive *entity.Archive) (*entity.Archive, error) {
	db := database.DB
	if err := db.Create(&archive).Error; err != nil {
		return nil, err
	}
	return archive, nil
}
func (*archive_repo) FindProductFromArchive(archive *entity.Archive, email, product_id string) error {
	db := database.DB
	if err := db.Where("email = ? AND product_id = ?", email, product_id).First(&archive).Error; err != nil {
		return err
	}
	return nil
}

func (*archive_repo) DeleteArchive(archive *entity.Archive, id uint) error {
	db := database.DB
	if err := db.Delete(&archive).Error; err != nil {
		return err
	}
	return nil
}
