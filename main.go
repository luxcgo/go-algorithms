package main

import "fmt"

func main() {
	d := &Dog{name: "a"}
	a := &Dog{name: "a", age: 4}
	fmt.Println(d == a)
}

type Dog struct {
	name string
	age  int
}

type Animal struct {
	name string
}
