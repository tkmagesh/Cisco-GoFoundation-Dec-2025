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

type PerishableProduct struct {
	Product
	Expiry string
}

// Overriding the Product.Format() method in PerishableProduct
func (pp PerishableProduct) Format() string {
	return fmt.Sprintf("%s, Expiry : %q", pp.Product.Format(), pp.Expiry)
}

func NewPerishableProduct(id int, name string, cost float64, expiry string) *PerishableProduct {
	return &PerishableProduct{
		Product: Product{
			Id:   id,
			Name: name,
			Cost: cost,
		},
		Expiry: expiry,
	}
}

func main() {

	milk := NewPerishableProduct(100, "Milk", 50, "2 Days")
	fmt.Println("Before applying discount :", milk.Format())
	milk.ApplyDiscount(10)
	fmt.Println("After applying discount :", milk.Format())
}
