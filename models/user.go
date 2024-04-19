package models

type User struct {
	Id       int    `json:"id" gorm:"primary_key"`
	Username string `json:"username"`
}

func (User) TableName() string {
	return "user"
}
