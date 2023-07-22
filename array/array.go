package main

import "fmt"

func main(){
	var name [3]string

	name[0] = "Valerian"
	name[1] = "Dwi"
	name[2] = "Purnomo"

	fmt.Println(name[0])
	fmt.Println(name[1])
	fmt.Println(name[2])

	var values = [3]int{
		80,
		80,
		10,
	}

	fmt.Println(len(values))
}