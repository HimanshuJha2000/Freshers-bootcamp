//Models/OrderModel.go
package Models

type Order struct{

	OrderID	  	uint	 	`json:"order_id" gorm:"primaryKey"`
	Customer    Customer 	`gorm:"foreignKey:CustomerID:"`
	CustomerID  uint     	`json:"customer_id" gorm:"default:null"`
	Product     Product  	`gorm:"foreignKey:ProductID:"`
	ProductID   uint     	`gorm:"default:null" json:"product_id"`
	Seller     	Seller  	`gorm:"foreignKey:SellerID:"`
	SellerID  	int	 		`json:"seller_id"`
	Quantity  	int    		`json:"quantity"`
	Status    	string 		`json:"status"`
}

func (b *Order) TableName() string {
	return "orders"
}