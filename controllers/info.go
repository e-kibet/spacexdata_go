package controllers

import (
	"demo/database"
	"demo/models"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type InfoRepo struct {
	Db *gorm.DB
}

func NewInfo() *InfoRepo {
	db := database.InitDb()
	db.AutoMigrate(&models.Info{})
	return &InfoRepo{Db: db}
}

func (repository *UserRepo) CreateInfo(c *gin.Context) {
	var info models.Info
	c.BindJSON(&info)
	err := models.CreateInfo(repository.Db, &info)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, info)
}

//get users
func (repository *InfoRepo) GetInfos(c *gin.Context) {
	var info []models.Info
	err := models.GetInfos(repository.Db, &info)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, info)
}

func (repository *InfoRepo) GetInfo(c *gin.Context) {
	id, _ := c.Params.Get("id")
	var info models.Info
	err := models.GetInfo(repository.Db, &info, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, info)
}

// update user
func (repository *InfoRepo) UpdateInfo(c *gin.Context) {
	var info models.Info
	id, _ := c.Params.Get("id")
	err := models.GetInfo(repository.Db, &info, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.BindJSON(&info)
	err = models.UpdateInfo(repository.Db, &info)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, info)
}

// delete user
func (repository *InfoRepo) DeleteInfo(c *gin.Context) {
	var info models.User
	id, _ := c.Params.Get("id")
	err := models.DeleteUser(repository.Db, &info, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
