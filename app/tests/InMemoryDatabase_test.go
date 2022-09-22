package tests

import (
	"FruitSale/app/Entities"
	"FruitSale/app/db"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetDiscounts(t *testing.T) {
	assert := assert.New(t)
	discountsTwo := db.DB.GetDiscounts()
	assert.Equal(len(discountsTwo), 0, "Returns zero records")
	db.DB.Seed()
	discounts := db.DB.GetDiscounts()
	assert.Equal(len(discounts), 2, "Returns only two discounts")
}

func TestGetProducts(t *testing.T) {
	db.DB.Seed()
	assert := assert.New(t)
	products := db.DB.GetProducts(1, 10)
	assert.Equal(len(products), 10, "Returns 10 products")
	db.DB.ClearData()
	productsTwo := db.DB.GetProducts(1, 10)
	assert.Equal(len(productsTwo), 0, "Returns zero products")
}

func TestGetProductLists(t *testing.T) {
	assert := assert.New(t)
	productsResponseTwo, err := db.DB.GetProductLists(1, 2)
	assert.Equal(len(productsResponseTwo.Products), 0, "Returns zero records")
	db.DB.Seed()
	productsResponse, err := db.DB.GetProductLists(1, 2)
	assert.Equal(len(productsResponse.Products), 2, "Returns only two records")
	assert.Equal(len(productsResponse.Discounts), 2, "Returns only two discounts")
	assert.Nil(err, "Error should be nil")
}

func TestGetProductByID(t *testing.T) {
	product := Entities.Product{
		ID:    100,
		Type:  "A100pple",
		Title: "Test title",
	}
	db.DB.AddProduct(product)
	assert := assert.New(t)
	product, err := db.DB.GetProductByID(100)
	assert.Equal(product.ID, 100, "Returns Product")
	assert.Nil(err, "Error should be nil")
	productResponseTwo, err := db.DB.GetProductByID(400000000)
	assert.Empty(productResponseTwo, "product should be empty")
	assert.NotNil(err)
	assert.Equal(err.Error(), "does not exist")
}

func TestAddProduct(t *testing.T) {
	db.DB.ClearData()
	product := Entities.Product{
		ID:    100,
		Type:  "Apple",
		Title: "Test title",
	}
	assert := assert.New(t)
	err := db.DB.AddProduct(product)
	assert.Nil(err, "Error should be nil for new product")
	duplicateProduct := db.DB.AddProduct(product)
	assert.NotNil(duplicateProduct, "Error should not be nil for new product with duplicate title")
	assert.Equal(duplicateProduct.Error(), "already exists")
}

func TestAUTO_ID(t *testing.T) {
	assert := assert.New(t)
	assert.True(db.AUTO_ID.ID() != db.AUTO_ID.ID(), true, "Autoincrement should uniq always")
}
