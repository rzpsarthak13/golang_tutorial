package Models

import "gorm.io/gorm"

type Customer struct {
	gorm.Model
	Name   string
	Orders []Order
}
