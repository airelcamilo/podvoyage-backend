package model

type UserResponse struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}
