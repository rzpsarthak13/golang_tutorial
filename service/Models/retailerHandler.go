package Models

import (
	"service/Config"

	_ "github.com/go-sql-driver/mysql"
)

func GetOrderByID(order *Order, id string) (err error) {
	if err := Config.DB.Where("id = ?", id).First(&order).Error; err != nil {
		return err
	}
	return nil
}

func GetOrderHistory(customerID string, order *[]Order) (err error) {
	if err := Config.DB.Where("customer_id = ?", customerID).Find(&order).Error; err != nil {
		return err
	}
	return nil
}

func GetTransactions(order *[]Order) (err error) {
	if err = Config.DB.Find(&order).Error; err != nil {
		return err
	}
	return nil
}

func GetPrice(product *Product, ProductID uint) (err error) {
	if err = Config.DB.First(&product, ProductID).Error; err != nil {
		return err
	}
	return nil
}
