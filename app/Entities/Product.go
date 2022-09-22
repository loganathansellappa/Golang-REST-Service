/*
Entities

	DB Schema
*/

package Entities

type Product struct {
	ID          int     `json:"id,"`
	Type        string  `json:"type"`
	ImageUrl    string  `json:"imageUrl"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
}

type ProductList struct {
	Products  []Product  `json:"products"`
	Discounts []Discount `json:"discounts"`
	Total     int        `json:"total"`
	limit     int        `json:"limit"`
	offset    int        `json:"offset"`
}

type TitleIndex struct {
	ProdId int
	Title  string
}
