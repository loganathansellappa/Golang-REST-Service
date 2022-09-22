/*
Entities
	DB Schema
*/

package Entities

var DiscountTypes = newDiscountTypes()
var ProductTypes = newProductTypes()

func newDiscountTypes() *types {
	return &types{
		Price:      "Price",
		Quantity:   "Quantity",
		Percentage: "Percentage",
	}
}

func newProductTypes() *FruitTypes {
	return &FruitTypes{
		Apple:     "Apple",
		Pineapple: "Pineapple",
		Banana:    "Banana",
		Others:    "Other",
	}
}

type types struct {
	Price      string
	Quantity   string
	Percentage string
}

type FruitTypes struct {
	Apple     string
	Pineapple string
	Banana    string
	Others    string
}

type Discount struct {
	ID          int    `json:"id"`
	Valid       bool   `json:"valid"`
	Type        string `json:"type"`
	ProductType string `json:"product_type"`
	Value       int    `json:"value"`
	Reward      int    `json:"reward"`
}
