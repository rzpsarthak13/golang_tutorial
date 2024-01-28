package Models

import (
	"service/Config"

	_ "github.com/go-sql-driver/mysql"
)

func CreateCustomer(customer *Customer) (err error) {
	if err = Config.DB.Create(customer).Error; err != nil {
		return err
	}
	return nil
}

func ValidateCustomer(customer *Customer, customerID uint) (err error) {
	if err := Config.DB.First(&customer, customerID).Error; err != nil {
		return err
	}
	return nil
}

func GetLastOrderTime(customer Customer, customerID uint) (err error) {
	if err := Config.DB.First(&customer, customerID).Error; err != nil {
		return err
	}
	return nil
}
