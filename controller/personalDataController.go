package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gitlab.com/pragmaticreviews/golang-gin-poc/entity"
	"gitlab.com/pragmaticreviews/golang-gin-poc/service"
)

type PersonalDataController interface {
	FindAll() []entity.PersonalData
	Save(ctx *gin.Context) entity.PersonalData
	FindById(ctx *gin.Context)
}

type controller struct {
	service service.PersonalDataService
}

func New(service service.PersonalDataService) PersonalDataController {
	return &controller{
		service: service,
	}
}

func (c *controller) FindAll() []entity.PersonalData {
	return c.service.FindAll()
}

func (c *controller) Save(ctx *gin.Context) entity.PersonalData {
	var data entity.PersonalData
	ctx.BindJSON(&data)
	c.service.Save(data)
	return data
}

// FindById implements PersonalDataController.
func (c *controller) FindById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Id must be a number",
		})
	}

	data, err := c.service.FindById(int(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Data tidak ditemukan"})
		return
	}

	ctx.JSON(http.StatusOK, data)
}
