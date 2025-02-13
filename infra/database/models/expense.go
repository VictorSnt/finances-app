package models

import "time"

type Expense struct {
	ID                 uint    `gorm:"primaryKey"`
	UserID             uint    `gorm:"not null;index"`
	Name               string  `gorm:"not null"`
	Amount             float64 `gorm:"not null"`
	ExpenseType        int     `gorm:"not null"`
	RecurrenceInMonths uint    `gorm:"not null"`
	CreatedAt          time.Time
}
