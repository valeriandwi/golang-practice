package main

import "fmt"

func EndApp () {
	fmt.Println("End App")
}

func RunApp(error bool) {
	defer EndApp()
	if error {
		panic("Error")
	}
}

func main(){
	RunApp(true)
}