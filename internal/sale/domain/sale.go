package domain

import (
	"time"

	"gorm.io/gorm"
)

type Sale struct {
	gorm.Model
	CustomerID uint
	Customer   Customer
	UserID     uint
	User       User
	Total      float64
	SaleDate   time.Time
	SaleItems  []SaleItem
}
