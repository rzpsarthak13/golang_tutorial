package Controllers

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"service/Models"
)

func PlaceOrder(c *gin.Context) {
	var order Models.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := validateProduct(order.ProductID, order.Quantity); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := validateCustomer(order.CustomerID); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// have to  check this once why this is not working????
	lastOrderTime, err := GetLastOrderTime(order.CustomerID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	cooldown := time.Since(lastOrderTime)
	if cooldown < 5*time.Minute {
		c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"error": "try again later"})
		return
	}

	order.Status = "order placed" // Default status

	if err := Models.PlaceOrder(&order); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if err := updateCustomerLastOrderTime(order.CustomerID, order.CreatedAt); err != nil {
		log.Println("Failed to update customer LastOrderTime:", err)
	}
	if err := UpdateProductQuantity(order.ProductID, order.Quantity); err != nil {
		log.Println("Failed to update order Quantity:", err)

	}
	c.JSON(http.StatusOK, order)
}

func validateProduct(productID uint, orderQuantity int) error {
	var product Models.Product
	if err := Models.ValidateProduct(&product, productID); err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.New("Invalid product ID")
		}
		return err
	}

	if product.Quantity < orderQuantity {
		return errors.New("insufficient product quantity")
	}

	return nil
}

func UpdateProductQuantity(productID uint, orderQuantity int) error {
	var product Models.Product
	product.Quantity -= orderQuantity
	if err := Models.UpdateProductQuantity(&product); err != nil {
		return errors.New("failed to update product quantity")
	}

	return nil
}
