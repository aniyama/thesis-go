package entities

type User struct {
	Id       uint
	Name     string `gorm:"size:16" json:"name"`
	Password string `gorm:"size:16" json:"password"`
}

type ReqUser struct {
	Name     string
	Password string
}
