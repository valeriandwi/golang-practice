package main

import "fmt"

func Ups() any {
	return "Ups"
}

func main() {
	empty := Ups()
	fmt.Println(empty)
}