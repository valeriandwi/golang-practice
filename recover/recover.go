package main

import "fmt"

func EndApp () {
	fmt.Println("End App")
	message := recover()
	fmt.Println("Panic ",message)
}

func RunApp(error bool) {
	defer EndApp()
	if error {
		panic("Error")
	}
}

func main(){
	RunApp(true)
	fmt.Println("Valerian")
}