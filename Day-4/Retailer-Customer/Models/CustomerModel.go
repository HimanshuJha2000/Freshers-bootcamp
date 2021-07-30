//Models/CustomerModel.go
package Models

type Customer struct {
	CustomerID      uint   	`json:"customer_id" gorm:"primaryKey"`
	CustomerName	string 	`json:"product_name"`
}

func (b *Customer) TableName() string {
	return "customers"
}