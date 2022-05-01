package models

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/abcd-edu/gentoo-socialgraph/internal/models/types"
)

func FollowUser(entry types.SocialGraph) error {
	sqlQuery := `
		INSERT INTO
		social_graph(follower_id, followed_id, created_on)
		VALUES($1,$2,$3)
		`

	_, err := db.Exec(sqlQuery, entry.FollowerId, entry.FollowedId, time.Now())
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func UnfollowUser(entry types.SocialGraph) error {
	sqlQuery := `
		DELETE FROM social_graph
		WHERE follower_id=$1 AND followed_id=$2
		`

	_, err := db.Exec(sqlQuery, entry.FollowerId, entry.FollowedId)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func GetFollowing(userId string) ([]types.UserSchema, error) {
	var followers []types.UserSchema
	sqlQuery := `
	SELECT *
	FROM social_graph
	WHERE follower_id=$1
	ORDER BY created_on 
	`

	rows, err := db.Query(sqlQuery, userId)
	if err != nil {
		fmt.Println(err)
		return followers, err
	}
	defer rows.Close()

	for rows.Next() {
		var followedId, followerId string
		var createdOn time.Time
		err := rows.Scan(&followerId, &followedId, &createdOn)
		if err != nil {
			fmt.Println(err)
			return followers, err
		}

		res, err := http.Get("http://localhost:8001/v1/user?user_id=" + userId)
		if err != nil {
			fmt.Println(err)
			return followers, err
		}
		defer res.Body.Close()

		userInfo := new(types.UserSchema)
		json.NewDecoder(res.Body).Decode(userInfo)

		followers = append(followers, *userInfo)
	}

	return followers, nil
}

func GetFollowers(userId string) ([]types.UserSchema, error) {
	var followers []types.UserSchema
	sqlQuery := `
	SELECT *
	FROM social_graph
	WHERE followed_id=$1
	ORDER BY created_on 
	`

	rows, err := db.Query(sqlQuery, userId)
	if err != nil {
		fmt.Println(err)
		return followers, err
	}
	defer rows.Close()

	for rows.Next() {
		var followedId, followerId string
		var createdOn time.Time
		err := rows.Scan(&followerId, &followedId, &createdOn)
		if err != nil {
			fmt.Println(err)
			return followers, err
		}

		res, err := http.Get("http://localhost:8001/v1/user?user_id=" + userId)
		if err != nil {
			fmt.Println(err)
			return followers, err
		}
		defer res.Body.Close()

		userInfo := new(types.UserSchema)
		json.NewDecoder(res.Body).Decode(userInfo)

		followers = append(followers, *userInfo)
	}

	return followers, nil
}
