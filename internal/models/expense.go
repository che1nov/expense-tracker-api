package models

type Expense struct {
	ID          uint `gorm:"primaryKey"`
	UserID      uint `gorm:"index"`
	Amount      float64
	Category    string
	Description string
	Date        string
}
