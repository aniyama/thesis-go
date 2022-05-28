package entity

type Tag struct {
	Id         int    `json:"id"`
	TagTitle   string `json:"title"`
	UserId     string `json:"userId"`
	CategoryId string `json:"categoryId"`
}

