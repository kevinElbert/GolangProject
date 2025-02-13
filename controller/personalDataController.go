package controller

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/pragmaticreviews/golang-gin-poc/entity"
	"gitlab.com/pragmaticreviews/golang-gin-poc/service"
)

type PersonalDataController interface {
	FindAll() []entity.PersonalData
	Save(ctx *gin.Context) entity.PersonalData
}

type controller struct {
	service service.PersonalDataService
}

func New(service service.PersonalDataService) PersonalDataController {
	return &controller {
		service: service,
	}
}

func (c *controller) FindAll() []entity.PersonalData{
	return c.service.FindAll()
}

func (c *controller) Save(ctx *gin.Context) entity.PersonalData{
	var data entity.PersonalData
	ctx.BindJSON(&data)
	c.service.Save(data)
	return data
}