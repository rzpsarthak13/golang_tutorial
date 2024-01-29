package Controllers

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"service/Config"
	"service/Models"
)

func CreateCustomer(c *gin.Context) {

	var customer Models.Customer
	c.BindJSON(&customer)
	customer.LastOrderTime = time.Now().Add(-10 * time.Minute)
	var err = Models.CreateCustomer(&customer)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, customer)
}

func validateCustomer(customerID uint) error {

	var customer Models.Customer
	if err := Models.ValidateCustomer(&customer, customerID); err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.New("Invalid customer ID")
		} else {
			return err
		}
	}
	return nil
}

func GetOrderHistory(c *gin.Context) {
	customerID := c.Params.ByName("id")
	var orders []Models.Order
	if err := Models.GetOrderHistory(customerID, &orders); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, orders)
}

func GetLastOrderTime(customerID uint) (time.Time, error) {
	var customer Models.Customer
	err := Models.GetLastOrderTime(&customer, customerID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return time.Time{}, errors.New("customer not found")
		}
	}
	return customer.LastOrderTime, nil
}

func updateCustomerLastOrderTime(customerID uint, newLastOrderTime time.Time) error {

	query := fmt.Sprintf("UPDATE customers SET last_order_time = ? WHERE id = ?")
	return Config.DB.Exec(query, newLastOrderTime, customerID).Error

}

//todo: delete customer func can also be added also special perks like for prime customer can also be done
