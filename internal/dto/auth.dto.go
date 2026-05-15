package dto

import "time"

type NewUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type User struct {
	Id        int        `json:"id"`
	Username  string     `json:"username"`
	Password  string     `json:"password,omitempty"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	Token     string     `json:"token,omitempty"`
}
