package domain

import "gorm.io/gorm"

type Supplier struct {
	gorm.Model
	Name          string
	ContactPerson string
	Email         string
	Phone         string
	Address       string
	Products      []Product
}
