package main

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/pragmaticreviews/golang-gin-poc/controller"
	"gitlab.com/pragmaticreviews/golang-gin-poc/service"
)

var (
	personalDataService    service.PersonalDataService       = service.New()
	personalDataController controller.PersonalDataController = controller.New(personalDataService)
)

func main() {
	r := gin.Default()

	// r.GET("/", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "Hellow World",
	// 	})
	// })
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, personalDataController.FindAll())
	})

	r.POST("/post", func(ctx *gin.Context) {
		ctx.JSON(200, personalDataController.Save(ctx))
	})
	r.Run(":8080")
}
