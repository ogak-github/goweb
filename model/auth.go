package model

type LoginRequest struct {
	Username string `json:"username" db:"username" validate:"required"`
	Password string `json:"password" db:"password" validate:"required"`
}

type RegisterUser struct {
	Username string `json:"username" db:"username" validate:"required"`
	Password string `json:"password" db:"password" validate:"required"`
	FullName string `json:"full_name" db:"full_name" validate:"required"`
	Email    string `json:"email" db:"email" validate:"email, required"`
}
