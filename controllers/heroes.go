package controllers

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"
	"tour-of-heroes-api-go/models"

	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
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

// PATCH /api/hero/:id
func UpdateHero(c *gin.Context) {
	// Get model if exist
	var hero models.Hero

	if err := models.DB.Where("id = ?", c.Param("id")).First(&hero).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Hero not found!"})
		return
	}

	// Validate input
	var input models.UpdateHeroInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&hero).Updates(input)

	c.JSON(http.StatusOK, hero)
}

// DELETE /api/hero/:id
func DeleteHero(c *gin.Context) {
	//Get model if exist
	var hero models.Hero

	if err := models.DB.Where("id = ?", c.Param("id")).First(&hero).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Hero not found!"})
		return
	}

	models.DB.Delete(&hero)

	c.JSON(http.StatusOK, true)
}

// GET /api/hero/alteregopic/:id
func GetAlterEgoPic(c *gin.Context) {

	var hero models.Hero

	if err := models.DB.Where("id = ?", c.Param("id")).First(&hero).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Hero not found!"})
		return
	}

	// Get image from Azure Storage //
	blobClient, err := azblob.NewClientFromConnectionString(os.Getenv("AZURE_STORAGE_CONNECTION_STRING"), nil)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Download the file
	pic, err := blobClient.DownloadStream(context.Background(), "alteregos", strings.ToLower(fmt.Sprintf("%s.png", strings.Replace(hero.AlterEgo, " ", "-", -1))), nil)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Send image to the browser
	c.DataFromReader(http.StatusOK, *pic.ContentLength, "image/png", pic.NewRetryReader(context.Background(), nil), nil)

}
