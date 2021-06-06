package api

import (
	"gin_example/app"
	"gin_example/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ReposIndex(c *gin.Context)  {
	var repos []model.Repo

	if err := app.DB.Find(&repos).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Record Not Found"})
		return
	}

	c.JSON(http.StatusOK, repos)
}

func ReposShow(c *gin.Context)  {
	id := c.Param("id")
	var repo model.Repo

	if err := app.DB.Where("id = ?", id).First(&repo).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Record Not Found"})
		return
	}

	c.JSON(http.StatusOK, repo)
}

func ReposCreate(c *gin.Context)  {
	var repo model.Repo

	if err := c.ShouldBindJSON(&repo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Params Invalid"})
		return
	}

	if err := app.DB.Create(&repo).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Create Failed"})
		return
	}

	c.JSON(http.StatusOK, repo)
}

func ReposUpdate(c *gin.Context)  {
	id := c.Param("id")
	var repo model.Repo

	if err := app.DB.Where("id = ?", id).Find(&repo).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Record Not Found"})
		return
	}

	if err := c.ShouldBindJSON(&repo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Params Invalid"})
		return
	}

	if err := app.DB.Updates(repo).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Update Failed"})
		return
	}

	c.JSON(http.StatusOK, repo)
}

func ReposDestroy(c *gin.Context)  {
	id := c.Param("id")
	var repo model.Repo

	if err := app.DB.Where("id = ?", id).Find(&repo).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Record Not Found"})
		return
	}

	if err := app.DB.Delete(&repo).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Destroy Failed"})
		return
	}

	c.JSON(http.StatusOK, repo)
}