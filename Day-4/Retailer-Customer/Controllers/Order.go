package Controllers

import (
	"first-api/Models"
	"github.com/gin-gonic/gin"
	"net/http"
)

//OrderProduct ... Order product for the user
func OrderProduct(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"message" : "Page not ready yet",
	})
}

func GetCustomerOrders(c *gin.Context) {
	var order []Models.Order
	id := c.Params.ByName("customer_id")
	err := Models.GetCustomerOrders(&order, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, order)
	}
}

func GetSellerTransactions(c *gin.Context) {
	var order []Models.Order
	id := c.Params.ByName("seller_id")
	err := Models.GetSellerTransactions(&order, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, order)
	}

}
