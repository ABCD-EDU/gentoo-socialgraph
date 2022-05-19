package models

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/abcd-edu/gentoo-socialgraph/internal/models/types"
)

func IsFollowing(userId string, followingId string) (bool, error) {
	sqlQuery := `
	SELECT follower_id FROM social_graph
	WHERE follower_id=$1 AND followed_id=$2
	`
	var check string
	if err := postDb.QueryRow(sqlQuery, userId, followingId).Scan(&check); err != nil {
		fmt.Println(err)
		return false, nil
	}

	return true, nil
}

func FollowUser(entry types.SocialGraph) error {
	sqlQuery := `
		INSERT INTO
		social_graph(follower_id, followed_id)
		VALUES($1,$2)
		`

	_, err := postDb.Exec(sqlQuery, entry.FollowerId, entry.FollowedId)
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

	_, err := postDb.Exec(sqlQuery, entry.FollowerId, entry.FollowedId)
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
	`

	rows, err := postDb.Query(sqlQuery, userId)
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
	`

	rows, err := postDb.Query(sqlQuery, userId)
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

func GetRandomUsers(userId string, amount int) ([]types.BasicUser, error) {
	var users []types.BasicUser
	sqlQuery := `
	SELECT *
	FROM users
	WHERE user_id NOT IN (
		SELECT social_graph.followed_id AS user_id
		FROM social_graph
		WHERE $1!=social_graph.follower_id
	) ORDER BY random() LIMIT $2;
	`

	rows, err := db.Query(sqlQuery, userId, amount)
	if err != nil {
		fmt.Println(err)
		return users, err
	}
	defer rows.Close()

	for rows.Next() {
		var userId, username, googlePhoto, email string
		err := rows.Scan(&userId, &username, &email, &googlePhoto)
		if err != nil {
			fmt.Println(err)
			return users, err
		}

		userToAdd := types.BasicUser{UserId: userId, Email: email, Username: username, GooglePhoto: googlePhoto}
		users = append(users, userToAdd)
	}

	return users, nil
}

func GetUserStat(column string, userId string) (int, error) {
	sqlQuery := fmt.Sprintf(`
		SELECT COUNT(*)
		FROM social_graph
		WHERE %s=$1;
	`, column)

	var stat int
	if err := postDb.QueryRow(sqlQuery, userId).Scan(&stat); err != nil {
		fmt.Println(err)
		return -1, err
	}

	return stat, nil
}

func containsUser(users []string, userId string) bool {
	for _, v := range users {
		if v == userId {
			return true
		}
	}
	return false
}

func SearchUser(query string, userId string, offset string, limit string) ([]types.SearchResult, error) {
	var users []types.BasicUser
	var followedIds []string
	var result []types.SearchResult
	sqlQuery := `
	SELECT user_id, username, email, google_photo, is_admin
	FROM users
	WHERE username LIKE '%' || $1 || '%' OR email LIKE '%' || $1 || '%' OR description LIKE '%' || $1 || '%'
	OFFSET $2 LIMIT $3;
	`
	rows, err := db.Query(sqlQuery, query, offset, limit)
	if err != nil {
		fmt.Println(err)
		return result, err
	}
	defer rows.Close()

	for rows.Next() {
		var userId, username, googlePhoto, email string
		var isAdmin bool
		err := rows.Scan(&userId, &username, &email, &googlePhoto, &isAdmin)
		if err != nil {
			fmt.Println(err)
			return result, err
		}

		userToAdd := types.BasicUser{UserId: userId, Email: email, Username: username, GooglePhoto: googlePhoto, IsAdmin: isAdmin}
		users = append(users, userToAdd)
	}

	followedQuery := `
	SELECT followed_id
	FROM social_graph
	WHERE follower_id=$1
	`

	following, err := postDb.Query(followedQuery, userId)
	if err != nil {
		return result, err
	}
	defer following.Close()

	for following.Next() {
		var followedId string
		err := following.Scan(&followedId)
		if err != nil {
			return result, err
		}

		followedIds = append(followedIds, followedId)
	}

	for _, v := range users {
		alreadyFollowed := containsUser(followedIds, v.UserId)
		resultToAppend := types.SearchResult{User: v, Followed: alreadyFollowed}
		result = append(result, resultToAppend)
	}

	return result, nil
}
