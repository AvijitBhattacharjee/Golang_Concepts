// Copyright (c) avijit bhattacharjee 2024

package Goroutine

import (
	"fmt"
)

// Define the cat struct
type cat struct {
	age int
}

// Define the dog struct
type dog struct {
	name string
}

// Define the animal interface
type animal interface {
	getName() string
	getAge() int
}

// Implement the getAge method for cat
func (c cat) getAge() int {
	return c.age
}

// Implement the getName method for cat
func (c cat) getName() string {
	return "Cat"
}

// Implement the getAge method for dog
func (d dog) getAge() int {
	return -1 // Dog does not have an age property, so return a default value
}

// Implement the getName method for dog
func (d dog) getName() string {
	return d.name
}

// ImplementInterface function
func ImplementInterface() {

	var a, b animal

	a = cat{age: 10}
	fmt.Println(a.getName(), "age:", a.getAge())

	b = dog{name: "tommy"}
	fmt.Println(b.getName(), "age:", b.getAge())
}
