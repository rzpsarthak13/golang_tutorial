package Models

import (
	"fmt"

	"service/Config"

	_ "github.com/go-sql-driver/mysql"
)

func CreateProduct(product *Product) (err error) {
	if err = Config.DB.Create(product).Error; err != nil {
		return err
	}
	return nil
}

func GetProduct(products *[]Product) (err error) {
	if err = Config.DB.Find(&products).Error; err != nil {
		return err
	}
	return nil
}

func GetProductByID(product *Product, id string) (err error) {
	if err = Config.DB.Where("id = ?", id).First(&product).Error; err != nil {
		return err
	}
	return nil
}

func UpdateProduct(product *Product, id string) (err error) {
	fmt.Println(product)
	Config.DB.Save(product)
	return nil
}

func DeleteProduct(product *Product, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(product)
	return nil
}

func ValidateProduct(product *Product, productID uint) (err error) {
	if err = Config.DB.First(&product, productID).Error; err != nil {
		return err
	}
	return nil
}

func UpdateProductQuantity(product *Product) (err error) {
	if err := Config.DB.Save(&product).Error; err != nil {
		return err
	}
	return nil
}
