package models

// User - User details
type User struct {
	Id       int32  `json:"id"`
	Username string `json:"username"`
	Hash     string `json:"-"`
}
