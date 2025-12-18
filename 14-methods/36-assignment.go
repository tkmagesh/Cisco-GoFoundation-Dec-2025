/*
Create a type for "ShoppingCart"
- should be able to add products to the cart and number of units for each product
- apply discount on the over all cart value
- print the cart (list the products and display the cart value)

Organize your code into modules & packages
*/

package main

import "fmt"

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

/*
Shopping Cart implementation
*/
func main() {
	// create an instance of the cart
	// add products and number of units
	// apply a discount to the whole cart
	// Print the cart items & cart value
}
