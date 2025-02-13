package service

import (
	"gitlab.com/pragmaticreviews/golang-gin-poc/entity"
)

type PersonalDataService interface {
	Save(entity.PersonalData) entity.PersonalData
	FindAll() []entity.PersonalData
}

type personalDataService struct {
	data []entity.PersonalData
}

func New() PersonalDataService {
	return &personalDataService{}
}

func (service *personalDataService) Save(datas entity.PersonalData) entity.PersonalData {
	service.data = append(service.data, datas)
	return datas
}

func (service *personalDataService) FindAll() []entity.PersonalData {
	return service.data
}
