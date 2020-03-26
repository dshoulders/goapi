package models

type List struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	UserId int32  `json:"userId"`
}
