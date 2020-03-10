package models

type ListItem struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Notes  string `json:"notes"`
	Owner  string `json:"owner"`
	UserId int    `json:"userId"`
}
