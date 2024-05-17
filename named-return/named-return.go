package main

import "fmt"

func getCompleteName() (firstName, middleName, lastName string) {
	firstName = "Valerian"
	middleName = "Dwi"
	lastName = "Purnomo"
	return firstName,middleName,lastName
}

func main(){
	a, b, c := getCompleteName()
	fmt.Println(a,b,c)
}