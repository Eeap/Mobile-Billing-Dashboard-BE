package models

// SignUp struct to describe register a new user.

// UserData struct to describe login user.
type UserData struct {
	Email    string `json:"email" validate:"required,email,lte=255"`
	Password string `json:"password" validate:"required,lte=255"`
}
