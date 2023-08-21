package models

type Basket struct {
	ID   uint  `gorm:"primaryKey"`
	Arts []Art `gorm:"many2many:basket_arts;"`
}
