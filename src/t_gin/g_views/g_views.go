package g_views

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"t_gin/models"
)

func Info(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "ok",
	})
}

func Post(c *gin.Context) {
	var DB *gorm.DB
	var users []models.User
	DB.Find(&users)
	c.JSON(http.StatusOK, gin.H{
		"info": users,
	})
}
