package main

import "fmt"

func getFullName() (string, string){
	return "Valerian", "Dwi Purnomo"
}

func main(){
	firstName, _ := getFullName()
	fmt.Println(firstName)
}