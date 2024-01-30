package main

import "fmt"

func getGoodbye(name string) string{
	return "Goodbye "+name
}

func main() {
	goodBye := getGoodbye

	result := goodBye("Valerian")

	fmt.Println(result)
}
