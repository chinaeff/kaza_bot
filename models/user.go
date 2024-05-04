package models

type User struct {
	UserID   int64   `json:"user_id"`
	Username string  `json:"username"`
	Orders   []Order `json:"orders"`
}
