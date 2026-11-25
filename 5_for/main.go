package main

import "fmt"

func main() {
	// while loop
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	// infinite loop
	// for {
	// 	fmt.Println("Hello, World")
	// }

	// classic for loop 
	for i := 1; i < 10; i++ {
		fmt.Println(i)
	}

	// ranage loop
	for i:= range 10 {
		fmt.Println(i)
	}
}