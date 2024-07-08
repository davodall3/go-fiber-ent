package model

type UserBody struct {
	Name     string  `json:"name" validate:"required"`
	Surname  string  `json:"surname" validate:"required"`
	Email    string  `json:"email" validate:"required,email"`
	Username string  `json:"username"`
	Password string  `json:"password"`
	Balance  float64 `json:"balance"`
}
