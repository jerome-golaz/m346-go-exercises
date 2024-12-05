package main

import "fmt"

func main() {
	// TODO: declare a type for Student (with first and last name)
	type Student struct {
		Firstname string
		Lastname  string
	}
	// TODO: declare a type for Class (consisting of multiple students)
	type Class struct {
		Students []Student
	}
	// TODO: declare a map of modules being attended by multiple classes
	modules := map[int]Class{
		114: {[]Student{{Firstname: "Hans", Lastname: "Peter"}, {Firstname: "Peter", Lastname: "Hans"}}},
		346: {[]Student{{Firstname: "Peter", Lastname: "Hans"}, {Firstname: "Hans", Lastname: "Peter"}}},
	}
	// TODO: output everything using fmt.Println()
	fmt.Println(modules)
}
