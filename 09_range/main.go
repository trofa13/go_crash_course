package main

import "fmt"

func main() {
	ids := []int{32, 33, 34, 35, 36}

	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	sum := 0
	for _, id := range ids {
		sum += id
	}

	fmt.Printf("Sum is %d\n", sum)

	phones := map[string]int{"Bob": 123, "Ivan": 456}
	for name, phone := range phones {
		fmt.Printf("%s phone is %d\n", name, phone)
	}
}
