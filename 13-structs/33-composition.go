package main

import "fmt"

type Product struct {
	Id   int
	Name string
	Cost float64
}

type PerishableProduct struct {
	Product
	Expiry string
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

	/*
		milk := PerishableProduct{
			Expiry: "2 days",
			Product: Product{
				Id:   100,
				Name: "Milk",
				Cost: 50,
			},
		}

		fmt.Println("Before applying discount :", Format(milk.Product))
		ApplyDiscount(&(milk.Product), 10)
		fmt.Println("After applying discount :", Format(milk.Product))
	*/

	milk := NewPerishableProduct(100, "Milk", 50, "2 Days")
	fmt.Println("Before applying discount :", Format(milk.Product))
	ApplyDiscount(&(milk.Product), 10)
	fmt.Println("After applying discount :", Format(milk.Product))
}

func Format(p Product) string {
	return fmt.Sprintf("Id : %d, Name : %q, Cost : %0.2f", p.Id, p.Name, p.Cost)
}

func ApplyDiscount(pPtr *Product, discountPercentage float64) { // no return result
	pPtr.Cost = pPtr.Cost * ((100 - discountPercentage) / 100)
}
