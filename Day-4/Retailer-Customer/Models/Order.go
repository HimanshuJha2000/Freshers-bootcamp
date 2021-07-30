package Models

import "first-api/Config"

func OrderProduct(order *Order) (err error){
	if err = Config.DB.Model(&order).Create(order).Error; err!=nil{
		return err
	}
	return nil
}

func GetCustomerOrders(order *[]Order, id string) (err error){
	if err = Config.DB.Model(&order).Where("customer_id = ?",id).Error; err!=nil{
		return err
	}
	return nil
}

func GetSellerTransactions(order *[]Order, id string) (err error) {
	if err = Config.DB.Find(&order).Where("seller_id = ?",id).Error; err != nil {
		return err
	}
	return nil
}
