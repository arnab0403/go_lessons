package main

import "fmt"

// we can declare variable in the outside of the function
const x = 10
func main() {
	const name = "Arnab"
	const age = 25

	fmt.Println(name, age)

	// constant grouping 

	const (
		PORT = 8080
		HOST = "localhost"
	)

	fmt.Println(PORT, HOST)

}