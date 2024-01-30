package main

import "fmt"

func main() {
	name := "Valerian"

	switch name {
	case "Valerian" :
			fmt.Println("Hello Valerian")
	case "Dwi" :
			fmt.Println("Hello Dwi")
	default :
		fmt.Println("Hiii")
	}

	length := len(name)
	switch{
	case length > 15 :
		fmt.Println("Nama terlalu panjang")
	case length > 5 :
		fmt.Println("Nama terlalu pendek")
	default :
		fmt.Println("Nama sudah benar")
	}
}