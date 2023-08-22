package models

type Art struct {
	ID          uint64 `gorm:"primaryKey"`
	Image       string `gorm:"not null"`
	Description string
	Name        string
	Price       float64
	UserId      uint64
	User        User     `gorm:"foreignKey:UserId"`
	Baskets     []Basket `gorm:"many2many:basket_arts;"`
}
