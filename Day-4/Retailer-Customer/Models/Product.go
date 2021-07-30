//Models/User.go
package Models
import (
	"first-api/Config"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

//GetAllProducts Fetch all products data
func GetAllProducts(product *[]Product) (err error) {
	if err = Config.DB.Find(product).Error; err != nil {
		return err
	}
	return nil
}

//CreateProduct ... Insert New data
func CreateProduct(product *Product) (err error) {
	if err = Config.DB.Create(product).Error; err != nil {
		return err
	}
	return nil
}

//GetProductByID ... Fetch only one product by Id
func GetProductByID(user *Product, id string) (err error) {
	if err = Config.DB.Where("product_id = ?", id).First(user).Error; err != nil {
		return err
	}
	return nil
}

//UpdateProduct ... Update product
func UpdateProduct(product *Product, id string) (err error) {
	fmt.Println(product)
	Config.DB.Save(product)
	return nil
}

//DeleteProduct ... Delete product
func DeleteProduct(product *Product, id string) (err error) {
	Config.DB.Where("product_id = ?", id).Delete(product)
	return nil
}