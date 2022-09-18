package entities

import "time"

type Module struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	UserId    int
	TagId     *int      `json:"tagId"` //0の時nilにする
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
