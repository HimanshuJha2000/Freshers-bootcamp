//Models/SellerModel.go
package Models

type Seller struct {
	SellerID    uint   	`json:"seller_id" gorm:"primaryKey"`
	SellerName	string 	`json:"seller_name"`
}

func (b *Seller) TableName() string {
	return "sellers"
}