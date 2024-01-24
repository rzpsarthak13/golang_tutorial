package Models

import (
	"time"

	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	Name          string
	LastOrderTime time.Time
}

type Order struct {
	gorm.Model
	CustomerID uint
	ProductID  uint
	Quantity   int
	Status     string
}

type Product struct {
	gorm.Model
	Name     string
	Quantity int
	Price    float64
}
