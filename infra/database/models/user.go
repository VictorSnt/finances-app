package models

type User struct {
	ID       uint    `gorm:"primaryKey"`
	Username string  `gorm:"unique;not null"`
	Income   float64 `gorm:"not null"`
}
