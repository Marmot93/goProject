package g_views

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"t_gin/models"
)

func Info(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "ok",
	})
}

func Post(c *gin.Context) {
	//var user models.User

	user := models.GetUsers()

	c.JSON(http.StatusOK, gin.H{
		"info": user,
	})
}
