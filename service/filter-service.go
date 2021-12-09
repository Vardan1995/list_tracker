package service

import (
	"github.com/Vardan1995/list_tracker/entity"
	"github.com/Vardan1995/list_tracker/repository"
)

var (
	filterRepository = repository.NewFilterRepository()
)

type filter_service struct{}

func NewFilterService() FilterService {
	return &filter_service{}
}

func (*filter_service) CreateFilter(filter *entity.Filter) (*entity.Filter, error) {
	if err := filterRepository.SaveFilter(filter); err != nil {
		return nil, err
	}
	return filter, nil

}

func (*filter_service) GetUserFilter(filters *[]entity.Filter, id uint) error {
	if err := filterRepository.FindFiltersByUserId(filters, id); err != nil {
		return err
	}
	return nil
}
func (*filter_service) GetFilters(filters *[]entity.Filter) error {
	if err := filterRepository.FindFilters(filters); err != nil {
		return err
	}
	return nil
}
