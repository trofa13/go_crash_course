package main

import "fmt"

func main() {
	emails := make(map[string]string)

	emails["Bob"] = "bob@gmail.com"
	emails["Sharon"] = "sharon@gmail.com"
	emails["Ivan"] = "ivan@gmail.com"

	fmt.Println(emails)
	fmt.Println(len(emails))
	fmt.Println(emails["Bob"])

	delete(emails, "Ivan")
	fmt.Println(emails)

	phones := map[string]int{"Bob": 123, "Ivan": 456}
	fmt.Println(phones)
}
