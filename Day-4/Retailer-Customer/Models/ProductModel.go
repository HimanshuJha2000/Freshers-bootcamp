//Models/ProductModel.go
package Models

type Product struct {
	ProductID   uint   	`json:"product_id" gorm:"primary_key"`
	ProdName    string 	`json:"product_name"`
	Price   	int 	`json:"price"`
	Quantity	int 	`json:"quantity"`
}

func (b *Product) TableName() string {
	return "products"
}