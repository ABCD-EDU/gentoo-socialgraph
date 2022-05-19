package services

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/abcd-edu/gentoo-socialgraph/internal/models"
	"github.com/abcd-edu/gentoo-socialgraph/internal/models/types"
	"github.com/gin-gonic/gin"
)

func FollowUser(c *gin.Context) {
	var query types.SocialGraph
	if err := c.ShouldBindJSON(&query); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "fail",
		})
		return
	}

	fmt.Println(query)
	c.Header("Content-Type", "application/json")
	err := models.FollowUser(query)
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
	var query types.SocialGraph
	if err := c.ShouldBindJSON(&query); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "fail",
		})
		return
	}

	fmt.Println(query)
	c.Header("Content-Type", "application/json")
	err := models.UnfollowUser(query)
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

func GetRandomUsers(c *gin.Context) {
	userId := c.Query("user_id")
	strAmount := c.Query("amount")
	intAmount, _ := strconv.Atoi(strAmount)

	users, err := models.GetRandomUsers(userId, intAmount)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "fail",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"users":  users,
	})
}

func GetUserStats(c *gin.Context) {
	userId := c.Query("user_id")

	following, err := models.GetUserStat("follower_id", userId)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "fail",
		})
		return
	}
	followers, err := models.GetUserStat("followed_id", userId)
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
		"following": following,
	})
}

func SearchUser(c *gin.Context) {
	userId := c.Query("user_id")
	query := c.Query("query")
	offset := c.Query("offset")
	limit := c.Query("limit")

	users, err := models.SearchUser(query, userId, offset, limit)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "fail",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"users":  users,
	})

}
