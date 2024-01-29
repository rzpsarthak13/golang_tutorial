package Routes

import (
	"github.com/gin-gonic/gin"

	"service/Controllers"
)

func SetUpRouter() *gin.Engine {
	// Product Management
	r := gin.Default()
	grp1 := r.Group("/products")
	{
		grp1.POST("", Controllers.CreateProduct)
		grp1.PATCH("/:id", Controllers.UpdateProduct)
		grp1.GET("/:id", Controllers.GetProductByID)
		grp1.GET("", Controllers.GetProduct)
		grp1.DELETE("/:id", Controllers.DeleteProduct)
	}
	grp2 := r.Group("/orders")
	{
		grp2.POST("", Controllers.PlaceOrder)
		grp2.GET("/:id", Controllers.GetOrderByID)
	}

	grp3 := r.Group("/customer")
	{
		grp3.POST("", Controllers.CreateCustomer)
		grp3.GET("/:id", Controllers.GetOrderHistory)
	}

	grp4 := r.Group("/retailer")
	{
		grp4.GET("", Controllers.GetTransactions)
	}

	return r
}
