package main

import "fmt"

func main() {
	// Arrays
	var fruitArr [2]string
	fruitArr[0] = "Apple"
	fruitArr[1] = "Banana"

	// declare and assign
	veggiesArr := [2]string{"Tomato", "Cucumber"}

	fmt.Println(fruitArr)
	fmt.Println(veggiesArr)

	// Slices
	fruitsSlice := []string{"Banana", "Apricot", "Pineaple", "Peach"}
	fmt.Println(fruitsSlice)
	fmt.Println(len(fruitsSlice))
	fmt.Println(fruitsSlice[1:3])

}
