package main

import "fmt"

func main() {
	a := 5
	b := &a

	fmt.Println("a value is", a)
	fmt.Println("b value is", b)
	fmt.Printf("a type is %T\n", a)
	fmt.Printf("b type is %T\n", b)
	fmt.Println("b pointer value is", *b)

	// change val woth pointer
	*b = 10
	fmt.Println("a now is", a)
}
