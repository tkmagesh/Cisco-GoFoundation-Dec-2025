package main

import "fmt"

type Product struct {
	Id   int
	Name string
	Cost float64
}

func main() {
	//var p Product // memory allocated & members initialized with their respective zero values
	// fmt.Println(p)
	// fmt.Printf("%#v\n", p)

	// var p Product = Product{Id: 100, Name: "Pen", Cost: 10}
	// var p = Product{Id: 100, Name: "Pen", Cost: 10}
	p := Product{Id: 100, Name: "Pen", Cost: 10}
	fmt.Printf("%+v\n", p)

	// Accessing the members
	fmt.Printf("Id : %d\n", p.Id)

	p2 := p // a copy will be created as structs are also values
	p2.Cost = 20
	fmt.Printf("p2.Cost = %f, p.Cost = %f\n", p2.Cost, p.Cost)

	// Using pointers for references
	var pPtr *Product
	pPtr = &p
	fmt.Printf("Id : %d, Name : %q, Cost : %f\n", pPtr.Id, pPtr.Name, pPtr.Cost)

	fmt.Println("Before applying discount, p :", Format( /* ? */ ))
	ApplyDiscount( /* ?? */ ) // apply 10% discount
	fmt.Println("After applying discount, p :", Format( /* ? */ ))
}

func Format( /* product */ ) string {
	return "Id : ?, Name : ?, Cost : ?"
}

func ApplyDiscount( /*  */ ) { // no return result
	// update the given products Cost by applying the given discount
}
