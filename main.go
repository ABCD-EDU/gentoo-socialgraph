package main

import (
	"github.com/abcd-edu/gentoo-socialgraph/internal/configs"
	"github.com/abcd-edu/gentoo-socialgraph/internal/models"
	"github.com/abcd-edu/gentoo-socialgraph/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	configs.InitializeViper()

	models.InitializeDB()
	models.InitializePostDB()

	router := gin.Default()

	router.Use(CORSMiddleware())
	v1 := router.Group("/v1")
	v1.Use(CORSMiddleware())
	{
		v1.GET("/", services.HandleMain)
		v1.GET("/suggested", services.GetRandomUsers)
		v1.GET("/followers", services.GetUserFollowers)
		v1.GET("/following", services.GetUserFollowing)
		v1.POST("/follow", services.FollowUser)
		v1.POST("/unfollow", services.UnfollowUser)
		v1.GET("/stats", services.GetUserStats)
		v1.GET("/search", services.SearchUser)
	}

	port := viper.GetString("serverPort")
	router.Run(":" + port)
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
