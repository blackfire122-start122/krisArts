package models

type User struct {
	Id       uint64 `gorm:"primaryKey"`
	Username string
	Password string
	Image    string
	Arts     []Art
	Basket   Basket `gorm:"foreignKey:BasketId"`
	BasketId uint64
}

type Admin struct {
	Id     uint64 `gorm:"primaryKey"`
	User   User   `gorm:"foreignKey:UserId"`
	UserId uint64
}
