package service

import (
	"gitlab.com/pragmaticreviews/golang-gin-poc/config"
	"gitlab.com/pragmaticreviews/golang-gin-poc/entity"
)

type PersonalDataService interface {
	Save(entity.PersonalData) entity.PersonalData
	FindAll() []entity.PersonalData
	FindById(id int) (entity.PersonalData, error)
}

type personalDataService struct {
	data []entity.PersonalData
}

func New() PersonalDataService {
	return &personalDataService{}
}

func (service *personalDataService) Save(data entity.PersonalData) entity.PersonalData {
	// service.data = append(service.data, datas)
	config.DB.Create(&data)
	return data
}

func (service *personalDataService) FindAll() []entity.PersonalData {
	var data []entity.PersonalData
	config.DB.Find(&data)
	return data
}

// FindById implements PersonalDataService.
func (service *personalDataService) FindById(id int) (entity.PersonalData, error) {
	var data entity.PersonalData
	result := config.DB.First(&data, id)
	if result.Error != nil{
		return entity.PersonalData{}, result.Error
	}
	return data, nil
}
