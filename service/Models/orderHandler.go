package Models

import (
	"service/Config"

	_ "github.com/go-sql-driver/mysql"
)

func PlaceOrder(order *Order) (err error) {
	if err = Config.DB.Create(&order).Error; err != nil {
		return err
	}
	return nil
}
