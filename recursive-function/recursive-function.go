package main

import "fmt"

func factorialRecursirve(value int) int {
	if(value == 1){
		return 1
	}else {
		return value * factorialRecursirve((value - 1))
	}
}

func main() {
	fmt.Println(factorialRecursirve((10)))
}