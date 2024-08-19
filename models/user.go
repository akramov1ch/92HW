package models

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type TrustedOrigin struct {
	UserID string `json:"user_id"`
	Origin string `json:"origin"`
}
