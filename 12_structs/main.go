package main

import (
	"fmt"
	"strconv"
)

// define a struct
type Person struct {
	firstName string
	lastName  string
	city      string
	gender    string
	age       int
}

// shorter alternative
// type Person struct {
// 	firstName, lastName, city, gender string
// 	age                               int
// }

// greeting method (value reciever)
func (person Person) greet() string {
	return "Hello, my name is " + person.firstName + " " + person.lastName + " and I am " + strconv.Itoa(person.age) + " years old"
}

// hasBirthday method (pointer reciever)
func (person *Person) hasBirthday() {
	person.age++
}

// getMarried (pointer reciever)
func (person *Person) getMarried(spouseLastname string) {
	if person.gender == "f" {
		person.lastName = spouseLastname
	}
}

func main() {
	person1 := Person{firstName: "Viktoria", lastName: "Minista", city: "LA", gender: "f", age: 25}

	// alternative
	// person1 := Person{"Viktor", "Minista", "LA", "f", 25}

	fmt.Println("hello mista", person1)
	fmt.Println("hello mista", person1.lastName)
	person1.age++
	fmt.Println("hello mista", person1.age)

	person1.hasBirthday()
	person1.getMarried("Williams")
	fmt.Println(person1.greet())
}
