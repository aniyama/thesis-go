package entities

type Tag struct {
	Id         int    `json:"id"`
	Title      string `json:"title"`
	UserId     int    `json:"userId"`
	CategoryId int    `json:"categoryId"`
}
