package model

type Session struct {
	Id    int `gorm:"primaryKey"`
	Token string
	User  User `gorm:"foreignKey:UserId;references:Id;"`
	UserId int 
}
