package model

type UserModel struct {
	ID       uint `gorm:"primary key"`
	Name     string
	Email    string `gorm:"unique"`
	Password string
	Status   string
}
