package Controllers

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"service/Config"
	"service/Models"
)

func CreateProduct(c *gin.Context) {
	var product Models.Product
	c.BindJSON(&product)
	var err = Config.DB.Create(&product).Error
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	} else {
		c.JSON(http.StatusOK, product)
	}
}

func GetProduct(c *gin.Context) {
	var products []Models.Product
	err := Config.DB.Find(&products).Error
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	} else {
		c.JSON(http.StatusOK, products)
	}
}

func GetProductByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var product Models.Product
	err := Config.DB.Where("id = ?", id).First(&product).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.AbortWithStatus(http.StatusNotFound)
		} else {
			c.AbortWithStatus(http.StatusInternalServerError)
		}
	} else {
		c.JSON(http.StatusOK, product)
	}
}

func UpdateProduct(c *gin.Context) {
	var product Models.Product
	id := c.Params.ByName("id")
	err := Config.DB.Where("id = ?", id).First(&product).Error // check if id exists
	if err != nil {
		c.JSON(http.StatusNotFound, product)
	}
	c.BindJSON(&product)
	err = Config.DB.Save(&product).Error
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	} else {
		c.JSON(http.StatusOK, product)
	}

}

func DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	var product Models.Product
	err := Config.DB.Where("id =?", id).Delete(&product).Error
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	} else {
		c.JSON(http.StatusOK, gin.H{"id": id, "message": "product deleted"})
	}
}

// CreateCustomer creates a new customer
func CreateCustomer(c *gin.Context) {
	var customerRequest struct {
		Name string `json:"name"`
		// Add any other fields you need for customer creation
	}

	if err := c.ShouldBindJSON(&customerRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create the customer
	newCustomer := Models.Customer{
		Name: customerRequest.Name,
		// Set other fields as needed
	}

	if err := Config.DB.Create(&newCustomer).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, newCustomer)
}

func PlaceOrder(c *gin.Context) {
	var order Models.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := validateProduct(order.ProductID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := validateCustomer(order.CustomerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	order.Status = "order placed" // Default status
	if err := Config.DB.Create(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, order)
}

func validateProduct(productID uint) error {
	var product Models.Product
	if err := Config.DB.First(&product, productID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.New("Invalid product ID")
		}
		return err
	}

	return nil
}

func validateCustomer(customerID uint) (*Models.Customer, error) {
	var customer Models.Customer
	result := Config.DB.First(&customer, customerID)
	if result.Error != nil {
		if errors.Is(result.Error, result.Error) {
			newCustomer := Models.Customer{
				Name: fmt.Sprintf("Customer_%d", customerID),
			}
			if err := Config.DB.Create(&newCustomer).Error; err != nil {
				return nil, err
			}
			return &newCustomer, nil
		}
		return nil, result.Error
	}

	return &customer, nil
}
