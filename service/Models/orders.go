package Models

import (
	"github.com/jinzhu/gorm"
)

type Order struct {
	gorm.Model
	CustomerID uint
	ProductID  uint
	Quantity   int
	Status     string
}
