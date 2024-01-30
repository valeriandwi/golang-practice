package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "Valerian",
		"address": "Jakarta",
	}

	person["title"] = "FE";

	fmt.Println(person);
	fmt.Println(person["name"]);
	fmt.Println(person["address"]);

	var book map[string]string = make(map[string]string)
	book["title"] = "Buku Golang"
	book["author"] = "Valerian"
	book["wrong"] = "Ups"

	delete(book,"wrong")

	fmt.Println(book)
}