package entity

type User struct {
	Id       uint
	Name     string `gorm:"size:16"`
	Password string `gorm:"size:16"`
	//   CreatedAt time.Time
}
