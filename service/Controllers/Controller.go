package Controllers

import (
	"errors"
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

func CreateCustomer(c *gin.Context) {

	var customer Models.Customer
	c.BindJSON(&customer)
	var err = Config.DB.Create(&customer).Error
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, customer)
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

	if err := validateCustomer(order.CustomerID); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//todo: order can be placed or failed.. by quantity availability
	//todo: get all transactions for a retailer

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

func validateCustomer(customerID uint) error {

	var customer Models.Customer
	if err := Config.DB.First(&customer, customerID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.New("Invalid customer ID")
		} else {
			return err
		}
	}
	return nil
}

func GetOrderByID(c *gin.Context) {
	var order Models.Order
	id := c.Params.ByName("id")
	err := Config.DB.Where("id = ?", id).First(&order).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.AbortWithStatus(http.StatusNotFound)
		} else {
			c.AbortWithStatus(http.StatusInternalServerError)
		}
	} else {
		c.JSON(http.StatusOK, order)
	}

}

func GetOrderHistory(c *gin.Context) {
	customerID := c.Params.ByName("id")
	var orders []Models.Order
	if err := Config.DB.Where("customer_id = ?", customerID).Find(&orders).Error; err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, orders)
}

func GetTransactions(c *gin.Context) {
	var orders []Models.Order
	err := Config.DB.Find(&orders).Error
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}

	var total float64
	for _, order := range orders {
		price, _ := GetPrice(order.ProductID)
		total += float64(order.Quantity) * price // here price has to be fetched by by product id
	}
	c.JSON(http.StatusOK, total)
}

func GetPrice(ProductID uint) (float64, error) {
	var product Models.Product
	err := Config.DB.First(&product, ProductID).Error
	if err != nil {
		return 0, errors.New("Product not found")
	}
	return product.Price, nil
}
