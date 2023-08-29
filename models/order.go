package models

type Order struct {
	ID     uint `gorm:"primaryKey"`
	Price  float64
	UserId uint64
	User   User  `gorm:"foreignKey:UserId"`
	Arts   []Art `gorm:"many2many:order_arts;"`
}
