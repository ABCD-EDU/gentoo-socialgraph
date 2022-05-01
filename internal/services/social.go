package services

import (
	"fmt"
	"net/http"
	"time"

	"github.com/abcd-edu/gentoo-socialgraph/internal/models"
	"github.com/abcd-edu/gentoo-socialgraph/internal/models/types"
	"github.com/gin-gonic/gin"
)

func FollowUser(c *gin.Context) {
	userId := c.Query("user_id")
	toFollow := c.Query("to_follow")

	c.Header("Content-Type", "application/json")
	err := models.FollowUser(types.SocialGraph{FollowerId: userId, FollowedId: toFollow, CreatedOn: time.Now()})
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "fail",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}

func UnfollowUser(c *gin.Context) {
	userId := c.Query("user_id")
	toFollow := c.Query("to_follow")

	c.Header("Content-Type", "application/json")
	err := models.UnfollowUser(types.SocialGraph{FollowerId: userId, FollowedId: toFollow, CreatedOn: time.Now()})
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "fail",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}

func GetUserFollowing(c *gin.Context) {
	userId := c.Query("user_id")

	c.Header("Content-Type", "application/json")
	following, err := models.GetFollowing(userId)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "fail",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":    "success",
		"following": following,
	})
}

func GetUserFollowers(c *gin.Context) {
	userId := c.Query("user_id")

	c.Header("Content-Type", "application/json")
	followers, err := models.GetFollowers(userId)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "fail",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":    "success",
		"followers": followers,
	})
}
