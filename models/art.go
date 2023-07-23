package models

type Art struct {
	ID          uint   `gorm:"primaryKey"`
	Image       string `gorm:"not null"`
	Description string
	Name        string
	Price       float64
	UserId      uint64
	User        User `gorm:"foreignKey:UserId"`
}
