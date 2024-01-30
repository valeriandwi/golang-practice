package main

import "fmt"

func main(){
	counter := 1

	for counter <= 10 {
		fmt.Println("Perulangan ke-",counter)
		counter++
	}

	for counter := 1; counter <= 10; counter ++ {
		fmt.Println("Perulangan ke-",counter)
	}

	slice := []string{"Valerian", "Dwi", "Purnomo"}

	for i:= 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	// with range
	for _, value:= range slice {
		fmt.Println(value)
	}

	person := make(map[string]string)
	person["name"] = "Valerian"
	person["title"] = "FE"

	for key, value:= range person {
		fmt.Println(key, "=", value)
	}
}