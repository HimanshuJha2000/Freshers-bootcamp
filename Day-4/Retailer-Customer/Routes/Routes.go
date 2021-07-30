//Routes/Routes.go
package Routes
import (
	"first-api/Controllers"
	"github.com/gin-gonic/gin"
)
//SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	r := gin.Default()
	grp1 := r.Group("/product")
	{
		grp1.GET("allProducts", Controllers.GetProducts)
		grp1.POST("createProduct", Controllers.CreateProduct)
		grp1.GET("products/:product_id", Controllers.GetProductByID)
		grp1.PATCH("updateProduct/:product_id", Controllers.UpdateProduct)
		grp1.DELETE("deleteProduct/:product_id", Controllers.DeleteProduct)
	}
	grp2 := r.Group("/customer")
	{
		grp2.POST("placeOrder",Controllers.OrderProduct)
		grp2.GET("getOrders/:customer_id", Controllers.GetCustomerOrders)

	}
	grp3 := r.Group("/seller")
	{
		grp3.GET("getTransactions/:seller_id", Controllers.GetSellerTransactions)
	}
	return r
}