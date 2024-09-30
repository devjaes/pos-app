package domain

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string
	Description string
	Price       float64
	Stock       int
	CategoryID  uint
	SupplierID  uint
}
