package main

import (
	"fmt"
	"os"
	"tour-of-heroes-api-go/controllers"
	"tour-of-heroes-api-go/models"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	// Load the .env file in the current directory
	godotenv.Load()

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to the Tour of Heroes API",
		})
	})

	router.GET("/api/hero", controllers.GetHeroes)
	router.GET("/api/hero/alteregopic/:id", controllers.GetAlterEgoPic)
	router.POST("/api/hero", controllers.CreateHero)
	router.GET("/api/hero/:id", controllers.FindHero)
	router.PATCH("/api/hero/:id", controllers.UpdateHero)
	router.DELETE("/api/hero/:id", controllers.DeleteHero)

	models.ConnectDatabase()

	port := os.Getenv("PORT")

	err := router.Run(fmt.Sprintf(":%v", port))

	if err != nil {
		return
	}
}
