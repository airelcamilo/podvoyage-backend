package model

type User struct {
	Id       int    `gorm:"primaryKey" json:"-"`
	Name     string `json:"name"`
	Email    string `gorm:"unique" json:"email"`
	Username string `gorm:"unique" json:"username"`
	Password []byte `json:"-"`
}
