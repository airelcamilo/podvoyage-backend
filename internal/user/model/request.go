package model

type RegisterRequest struct {
	Name     string
	Email    string
	Username string
	Password string
}

type LoginRequest struct {
	Email    string
	Password string
}

type ValidateRequest struct {
	Token string
}
