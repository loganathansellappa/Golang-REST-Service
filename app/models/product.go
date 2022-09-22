/*
Model
	Responsible for handling data
*/

package models

import (
	Entities "FruitSale/app/Entities"
	"FruitSale/app/db"
)

func ListAllProducts(offset int, limit int) (*Entities.ProductList, error) {
	productsResponse, err := db.DB.GetProductLists(offset, limit)
	if err != nil {
		return nil, err
	}
	return productsResponse, nil
}

func ListProduct(id int) (Entities.Product, error) {
	product, err := db.DB.GetProductByID(id)
	if err != nil {
		return Entities.Product{}, err
	}
	return product, nil
}

func UpdateProduct(prod Entities.Product) (*Entities.Product, error) {
	product, err := db.DB.UpdateProduct(prod.Title, prod.ID)
	if err != nil {
		return nil, err
	}
	return product, nil
}
