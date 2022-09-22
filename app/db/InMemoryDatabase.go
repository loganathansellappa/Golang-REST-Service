/*
MemoryDatabase
	Inmemory storage to store the records in maps
*/

package db

import (
	"errors"
	"fmt"
	"sort"
	"sync"
)

import "FruitSale/app/Entities"

var DB *MemoryDatabase = NewMemoryDatabase()
var AUTO_ID autoIncrement
var (
	ErrDoesNotExist  = errors.New("does not exist")
	ErrAlreadyExists = errors.New("already exists")
)

/*
autoIncrement

	Generates uniq numbers
*/
type autoIncrement struct {
	sync.Mutex // ensures autoInc is goroutine-safe
	id         int
}

func (a *autoIncrement) ID() (id int) {
	a.Lock()
	defer a.Unlock()

	id = a.id

	a.id++
	return
}

/*
MemoryDatabase

	Database schema to store the entities Product/Discount/Indices
*/
type MemoryDatabase struct {
	lock      sync.RWMutex
	products  map[string]Entities.Product
	discounts map[string]Entities.Discount
	indices   map[string]Entities.TitleIndex
}

/*
NewMemoryDatabase

	Create inmemory store using maps
*/
func NewMemoryDatabase() *MemoryDatabase {
	return &MemoryDatabase{
		products:  make(map[string]Entities.Product),
		discounts: make(map[string]Entities.Discount),
		indices:   make(map[string]Entities.TitleIndex),
	}
}

/*
GetDiscounts

	Reads Discounts data from store
*/
func (d *MemoryDatabase) GetDiscounts() []Entities.Discount {
	d.lock.RLock()
	defer d.lock.RUnlock()
	discounts := make([]Entities.Discount, 0, len(d.discounts))
	for _, value := range d.discounts {
		discounts = append(discounts, value)
	}
	sort.Slice(discounts, func(i, j int) bool {
		return discounts[i].ID < discounts[j].ID
	})

	return discounts
}

/*
GetProducts

	Reads Products data from store
*/
func (d *MemoryDatabase) GetProducts(offset int, limit int) []Entities.Product {
	d.lock.RLock()
	defer d.lock.RUnlock()

	// Make a copy of the products map (as a slice)
	products := make([]Entities.Product, 0, limit)
	for _, product := range d.products {
		if product.ID >= offset && product.ID < offset+limit {
			products = append(products, product)
		}
	}
	// Sort by ID so we return them in a defined order
	sort.Slice(products, func(i, j int) bool {
		return products[i].ID < products[j].ID
	})
	return products
}

/*
GetProductLists

	Formatted response for list requset
*/
func (d *MemoryDatabase) GetProductLists(offset int, limit int) (*Entities.ProductList, error) {
	d.lock.RLock()
	defer d.lock.RUnlock()
	productsResponse := new(Entities.ProductList)
	productsResponse.Products = d.GetProducts(offset, limit)
	productsResponse.Discounts = d.GetDiscounts()
	productsResponse.Total = len(d.products)
	return productsResponse, nil
}

/*
GetProductByID

	Fetch the product by id
*/
func (d *MemoryDatabase) GetProductByID(id int) (Entities.Product, error) {
	d.lock.RLock()
	defer d.lock.RUnlock()
	product, ok := d.products["p"+string(id)]
	if !ok {
		return Entities.Product{}, ErrDoesNotExist
	}
	return product, nil
}

/*
AddProduct

	Adds new product to the store
*/
func (d *MemoryDatabase) AddProduct(product Entities.Product) error {
	d.lock.Lock()
	defer d.lock.Unlock()

	if _, ok := d.products["p"+string(product.ID)]; ok {
		return ErrAlreadyExists
	}
	d.products["p"+string(product.ID)] = product
	return nil
}

/*
AddDiscount

	Adds new discount to the store
*/
func (d *MemoryDatabase) AddDiscount(discount Entities.Discount) error {
	d.lock.Lock()
	defer d.lock.Unlock()

	if _, ok := d.discounts[string(discount.ID)]; ok {
		return ErrAlreadyExists
	}
	d.discounts[string(discount.ID)] = discount
	return nil
}

/*
AddTitleIndex

	Adds index for title to speed up the lookup
*/
func (d *MemoryDatabase) AddTitleIndex(product Entities.Product) error {
	d.lock.Lock()
	defer d.lock.Unlock()

	if _, ok := d.indices[product.Title]; ok {
		return ErrAlreadyExists
	}
	d.indices[product.Title] = Entities.TitleIndex{ProdId: product.ID, Title: product.Title}
	return nil
}

/*
RemoveTitleIndex

	Removes invalid indices
*/
func (d *MemoryDatabase) RemoveTitleIndex(indice string) {
	d.lock.Lock()
	defer d.lock.Unlock()
	delete(d.indices, indice)
}

/*
UpdateProduct

	Updates the title of the product
*/
func (d *MemoryDatabase) UpdateProduct(title string, id int) (*Entities.Product, error) {
	product, ok := d.products["p"+string(id)]
	if !ok {
		return nil, ErrDoesNotExist
	}

	titleIndex, validTitle := d.indices[title]

	if !validTitle || titleIndex.ProdId == id {
		d.RemoveTitleIndex(product.Title)
		product.Title = title
		d.AddTitleIndex(product)
		d.products["p"+string(product.ID)] = product
		return &product, nil
	}
	return nil, ErrAlreadyExists

}

/*
Seed

	Seed dummy data foe the app
*/
func (d *MemoryDatabase) Seed() {
	for i := 1; i <= 30; i++ {
		d.SeedProduct(AUTO_ID.ID(), Entities.ProductTypes.Apple, 1)
		d.SeedProduct(AUTO_ID.ID(), Entities.ProductTypes.Banana, 2)
		d.SeedProduct(AUTO_ID.ID(), Entities.ProductTypes.Pineapple, 5)
	}
	d.SeedDiscount()
}
func (d *MemoryDatabase) SeedDiscount() {
	discount := Entities.Discount{
		ID:          AUTO_ID.ID(),
		Valid:       true,
		Type:        Entities.DiscountTypes.Quantity,
		ProductType: Entities.ProductTypes.Apple,
		Value:       2,
		Reward:      1,
	}
	d.AddDiscount(discount)

	discountTwo := Entities.Discount{
		ID:          AUTO_ID.ID(),
		Valid:       true,
		Type:        Entities.DiscountTypes.Price,
		ProductType: Entities.ProductTypes.Pineapple,
		Value:       2,
		Reward:      1,
	}
	d.AddDiscount(discountTwo)
}
func (d *MemoryDatabase) SeedProduct(id int, name string, price float32) {
	title := fmt.Sprintf(`%s - %d`, name, id)
	prod := Entities.Product{
		ID:          id,
		Type:        name,
		Title:       title,
		Description: name + "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed bibendum ante purus, id fermentum nulla commodo vehicula. Nulla id volutpat lorem, et vehicula enim. Nam erat nunc, efficitur eget nibh at, euismod molestie justo. Nunc dictum euismod lacus nec varius. Sed efficitur dui et magna elementum pharetra. Proin ipsum mauris, pretium id tellus at, ultrices maximus metus. Interdum et malesuada fames ac ante ipsum primis in faucibus. In arcu dui, volutpat in sem vehicula, hendrerit lacinia elit. Nullam tristique augue non tempor efficitur. Nulla iaculis nibh at tortor pharetra, sit amet ornare lectus auctor. Sed nec sapien at est fermentum posuere. Fusce suscipit tristique velit, et consectetur nunc dignissim nec. Sed id ligula ut ante gravida ornare. Suspendisse eu odio lacus. Proin rhoncus convallis vehicula.",
		Price:       price,
	}
	d.AddProduct(prod)
	d.AddTitleIndex(prod)
}
