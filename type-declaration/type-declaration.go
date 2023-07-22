package main

import "fmt"

func main(){
	type noKTP string
	type married bool

	var noKTPVale noKTP = "01100101"
	var marriedStatus married = false

	fmt.Println(noKTPVale)
	fmt.Println(marriedStatus)
}