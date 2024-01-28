package Controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"service/Models"
)

func GetOrderByID(c *gin.Context) {
	var order Models.Order
	id := c.Params.ByName("id")
	err := Models.GetOrderByID(&order, id)
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

func GetTransactions(c *gin.Context) {
	var orders []Models.Order
	err := Models.GetTransactions(&orders)
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
	err := Models.GetPrice(&product, ProductID)
	if err != nil {
		return 0, errors.New("Product not found")
	}
	return product.Price, nil
}
