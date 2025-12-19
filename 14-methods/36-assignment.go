/*
Create a type for "ShoppingCart"
- should be able to add products to the cart and number of units for each product
- apply discount on the over all cart value
- print the cart (list the products and display the cart value)

Organize your code into modules & packages
*/

package main

import (
	"fmt"
	"strings"
)

type Product struct {
	Id   int
	Name string
	Cost float64
}

func (p Product) Format() string {
	return fmt.Sprintf("Id : %d, Name : %q, Cost : %0.2f", p.Id, p.Name, p.Cost)
}

func (pPtr *Product) ApplyDiscount(discountPercentage float64) { // no return result
	pPtr.Cost = pPtr.Cost * ((100 - discountPercentage) / 100)
}

func NewProduct(id int, name string, cost float64) *Product {
	return &Product{
		Id:   id,
		Name: name,
		Cost: cost,
	}
}

type LineItem struct {
	Item  *Product
	Units float64
}

func (li *LineItem) ItemValue() float64 {
	return li.Units * li.Item.Cost
}

func (li LineItem) Format() string {
	return fmt.Sprintf("%s, Units : %0.2f", li.Item.Format(), li.Units)
}

func NewLineItem(product *Product, units float64) *LineItem {
	return &LineItem{
		Item:  product,
		Units: units,
	}
}

type ShoppingCart struct {
	LineItems []*LineItem
	Discount  float64
}

func (cart *ShoppingCart) AddItem(p *Product, units float64) {
	li := NewLineItem(p, units)
	cart.LineItems = append(cart.LineItems, li)
}

func (cart ShoppingCart) CartValue() float64 {
	var total float64
	for _, li := range cart.LineItems {
		total += li.ItemValue()
	}
	// apply the discount
	total = total * ((100 - cart.Discount) / 100)
	return total
}

func (cart ShoppingCart) Format() string {
	var sb strings.Builder
	sb.WriteString("My Shopping Cart\n")
	sb.WriteString("================================================\n")
	for _, li := range cart.LineItems {
		sb.WriteString(li.Format())
		sb.WriteString("\n")
	}
	sb.WriteString(fmt.Sprintf("Discount : %0.2f\n", cart.Discount))
	sb.WriteString(fmt.Sprintf("Total : %0.2f\n", cart.CartValue()))
	return sb.String()
}

func (cart *ShoppingCart) ApplyDiscount(discountPercentage float64) {
	cart.Discount = discountPercentage
}

func NewShoppingCart() *ShoppingCart {
	// return &ShoppingCart{}
	return new(ShoppingCart)
}

// in-memory product repo
var products map[string]*Product = map[string]*Product{
	"pen":    NewProduct(100, "Pen", 10),
	"pencil": NewProduct(101, "Pencil", 5),
	"marker": NewProduct(102, "Marker", 50),
}

/*
Shopping Cart implementation
*/
func main() {

	// create an instance of the cart
	cart := NewShoppingCart()

	// add products and number of units
	cart.AddItem(products["pen"], 10)
	cart.AddItem(products["pencil"], 20)
	cart.AddItem(products["marker"], 5)

	// apply a discount to the whole cart
	cart.ApplyDiscount(5)

	// Print the cart items & cart value
	fmt.Println(cart.Format())
}
