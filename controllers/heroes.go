package controllers

import (
	"tour-of-heroes-api-go/models"

	"github.com/gin-gonic/gin"
)

// GET /api/hero
func GetHeroes(c *gin.Context) {
	var heroes []models.Hero
	models.DB.Find(&heroes)

	c.JSON(200, heroes)
}
