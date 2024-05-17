package main

import "fmt"

type Customer struct {
	Name    string
	Address string
	Age     int
}

func (customer Customer) sayHello(name string){
	fmt.Println("Hello ",name, "my name is", customer.Name)
}

func main() {
	var customer Customer
	customer.Name = "Valerian"
	customer.Address = "Indonesia"
	customer.Age = 27

	fmt.Println(customer)

	joko := Customer {
		Name : "Joko",
		Address : "Indonesia",
		Age : 30,
	}
	fmt.Println(joko)

	budi := Customer{"Budi","Indonesia",20}
	fmt.Println(budi)

	budi.sayHello("Agus")
	
}