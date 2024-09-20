package domain

import (
	"gorm.io/gorm"
)

type SaleItem struct {
	gorm.Model
	SaleID    uint
	Sale      Sale
	ProductID uint
	Product   Product
	Quantity  int
	UnitPrice float64
	Subtotal  float64
}
