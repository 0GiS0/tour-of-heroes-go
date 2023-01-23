package main

import (
	"tour-of-heroes-api-go/controllers"
	"tour-of-heroes-api-go/models"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to the Tour of Heroes API",
		})
	})

	router.GET("/api/hero", controllers.GetHeroes)
	router.POST("/api/hero", controllers.CreateHero)
	router.GET("/api/hero/:id", controllers.FindHero)
	router.PATCH("/api/hero/:id", controllers.UpdateHero)

	models.ConnectDatabase()

	router.Run("localhost:8080")
}
