package Controllers

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

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
