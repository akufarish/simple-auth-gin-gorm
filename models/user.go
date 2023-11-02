package models

type User struct {
	Id int `gorm:"primaryKey" json:"id"`
	Name string `gorm:"type:varchar(255)" json:"name"`
	Email string `gorm:"type:varchar(255)" json:"email"`
	Password string `gorm:"type:varchar(255)" json:"password"`
}