package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleMain(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"api": "Social Graph Service",
	})
}
