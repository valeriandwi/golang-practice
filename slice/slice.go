package main

import "fmt"

func main(){
	var months = [12]string {
		"January",
		"February",
		"March",
		"April",
		"May",
		"June",
		"July",
		"August",
		"September",
		"October",
		"November",
		"December",
	}

	var slice1 = months[4:7]
	fmt.Println(slice1);

	var slice2 = months[10:]
	fmt.Println(slice2)

	var slice3 = append(slice2,"Valerian")
	fmt.Println(slice3)
	slice3[1] = "Not December"
	fmt.Println(slice3)

	newSlice := make([]string,2,5)
	newSlice[0] = "Valerian"
	newSlice[1] = "Dwi"

	fmt.Println(newSlice)

	copySlice := make([]string,len(newSlice),cap(newSlice))
	copy(copySlice,newSlice)
	fmt.Println(copySlice)

	iniArray := [5]int{1,2,3,4,5}
	iniSlice := []int{1,2,3,4,5}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}