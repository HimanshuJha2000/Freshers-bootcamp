//Controllers/User.go
package Controllers
import (
	"first-api/Models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

//GetProducts ... Get all Products
func GetProducts(c *gin.Context) {
	var product []Models.Product
	err := Models.GetAllProducts(&product)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, product)
	}
}

//CreateProduct ... Create product
func CreateProduct(c *gin.Context) {
	var product Models.Product
	c.BindJSON(&product)
	err := Models.CreateProduct(&product)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"id": product.ProductID,
			"product_name": product.ProdName,
			"price": product.Price,
			"quantity": product.Quantity,
			"message": "product successfully added",
		})
	}
}

//GetProductByID ... Get the product by id
func GetProductByID(c *gin.Context) {
	id := c.Params.ByName("product_id")
	var user Models.Product
	err := Models.GetProductByID(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

//UpdateProduct ... Update the product information
func UpdateProduct(c *gin.Context) {
	var product Models.Product
	id := c.Params.ByName("product_id")
	err := Models.GetProductByID(&product, id)
	if err != nil {
		c.JSON(http.StatusNotFound, product)
	}
	c.BindJSON(&product)
	err = Models.UpdateProduct(&product, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, product)
	}
}

//DeleteProduct ... Delete the product
func DeleteProduct(c *gin.Context) {
	var product Models.Product
	id := c.Params.ByName("product_id")
	err := Models.DeleteProduct(&product, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}
}