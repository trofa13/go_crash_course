package main

import "fmt"

func main() {
	x := 5
	y := 10

	if x < y {
		fmt.Printf("%d is less than %d\n", x, y)
	} else {
		fmt.Printf("%d is less than %d\n", y, x)
	}

	color := "red"

	if color == "red" {
		fmt.Println("color is red")
	} else if color == "blue" {
		fmt.Println("color is NOT blue or red")
	} else {
		fmt.Println("color is NOT blue or red")
	}

	switch color {
	case "red":
		fmt.Println("color is red")
	case "blue":
		fmt.Println("color is NOT blue or red")
	default:
		fmt.Println("color is NOT blue or red")
	}
}
