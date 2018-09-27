package main

import "fmt"

func main() {
	// MAIN TYPES
	// string
	// bool
	// int int8 int16 int32 int64
	// uint uint8 uint16 uint32 uint64 uintptr
	// byte - alias for uint8
	// rune - alias for int32
	// float32 float64
	// complex64 complex128k

	// using var
	var name string = "Trofim"
	fmt.Println(name)
	fmt.Printf("%T\n", name)

	// using const
	const age int32 = 25
	// ! age = 26

	// shothand
	surname := "Samusev"
	fmt.Println(surname)

	login, email := "trofa13", "trofa13@gmail.com"
	fmt.Println(login, email)

}
