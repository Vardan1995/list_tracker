package service

import (
	"github.com/Vardan1995/filter_notifyer/entity"
	"github.com/Vardan1995/filter_notifyer/repository"
)

type archive_service struct{}

var (
	archiveRepository = repository.NewArchiveRepository()
)

func NewArchiveService() ArchiveService {
	return &archive_service{}
}

func (*archive_service) CreateIfNotExists(email, product_id string) bool {
	var archive entity.Archive

	if err := archiveRepository.FindProductFromArchive(&archive, email, product_id); err != nil {
		archive.Email = email
		archive.ProductId = product_id
		archiveRepository.SaveArchive(&archive)
		return false
	}

	return true
}
