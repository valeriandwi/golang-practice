package main

import "fmt"

func getHello(name string) string {
	if(name == ""){
		return "Kosong"
	} else {
		return "Hello"
	}
}

func main() {
	result := getHello("Valerian")

	fmt.Println(result)
	fmt.Println(getHello(""))
}