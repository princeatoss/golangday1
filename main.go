package main

import (
	"fmt"
	"learngo/workshop"
)

func main() {
	fmt.Println("test")
	var name int
	fmt.Scanln(&name)
	a, _ := add(1, 2)

	fmt.Println(a)
	workshop.Print("test")
	for test := 0; test < 10; test++ {
		fmt.Printf("I am number %v\n", test)
	}
	//var people [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	people := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	for i := 0; i < len(people); i++ {
		fmt.Printf("I am person %v\n", people[i])
	}
	for i, _ := range people {
		fmt.Printf("I am person %v\n", people[i])
	}
	var newIntern = intern{
		name:   "name",
		gender: "Male",
		age:    name,
	}
	fmt.Print(newIntern)
}

func add(a int, b int) (int, int) {
	return a + b, a - b
}

type intern struct {
	name   string
	gender string
	age    int
}
