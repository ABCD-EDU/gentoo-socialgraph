package types

import "time"

type SocialGraph struct {
	FollowerId string    `form:"follower_id" json:"follower_id" xml:"follower_id"  binding:"required"`
	FollowedId string    `form:"followed_id" json:"followed_id" xml:"followed_id"  binding:"required"`
	CreatedOn  time.Time `form:"created_on" json:"created_on" xml:"created_on"  binding:"required"`
}

type UserSchema struct {
	UserId    string           `form:"user_id" json:"user_id" xml:"user_id"  binding:"required"`
	UserInfo  UserRegistration `form:"user_info" json:"user_info" xml:"user_info"  binding:"required"`
	CreatedOn time.Time        `form:"created_on" json:"created_on" xml:"created_on"  binding:"required"`
	CanPost   bool             `form:"can_post" json:"can_post" xml:"can_post"  binding:"required"`
}

type UserRegistration struct {
	Email       string `form:"email" json:"email" xml:"email"  binding:"required"`
	Username    string `form:"username" json:"username" xml:"username"  binding:"required"`
	GooglePhoto string `form:"google_photo" json:"google_photo" xml:"google_photo"  binding:"required"`
	Description string `form:"description" json:"description" xml:"description"  binding:"required"`
}
