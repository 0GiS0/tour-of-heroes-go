package controllers

import (
	"net/http"
	"tour-of-heroes-api-go/models"

	"github.com/gin-gonic/gin"
)

// GET /api/hero
func GetHeroes(c *gin.Context) {
	var heroes []models.Hero
	models.DB.Find(&heroes)

	c.JSON(200, heroes)
}

// POST /api/hero
func CreateHero(c *gin.Context) {
	// Validate input
	var input models.CreateHeroInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create hero
	hero := models.Hero{Name: input.Name, AlterEgo: input.AlterEgo, Description: input.Description}
	models.DB.Create(&hero)

	c.JSON(http.StatusOK, hero)
}

// GET /api/hero/:id
func FindHero(c *gin.Context) {
	var hero models.Hero

	if err := models.DB.Where("id = ?", c.Param("id")).First(&hero).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Hero not found!"})
		return
	}

	c.JSON(http.StatusOK, hero)
}
