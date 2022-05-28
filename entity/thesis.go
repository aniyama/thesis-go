package entity

import "time"

type Thesis struct {
	Id            int    `json:"id"`
	ThesisTitle   string `json:"thesisTitle"`
	ThesisContent string `json:"thesisContent"`
	UserId        int
	TagId         *int    `json:"tagId"`//0の時nilにする
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
