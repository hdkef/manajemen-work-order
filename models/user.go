package models

type User struct {
	ID       int64  `json:"user_id"`
	FullName string `json:"user_full_name"`
	Username string `json:"user_username"`
	Password string `json:"user_password"`
	Role     string `json:"user_role"`
}
