package model

import "time"

type Todo struct {
	Id        string    `json:"id" db:"id"`
	Title     string    `json:"title" db:"title"`
	Content   string    `json:"content" db:"content" validate:"required"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	ModifyAt  time.Time `json:"modify_at" db:"modify_at"`
	UserId    string    `json:"user_id" db:"user_id" validate:"required"`
}
