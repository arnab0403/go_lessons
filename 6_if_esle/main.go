package main

import "fmt"

func main () {
	// age := 25 

	// if age > 18 {
	// 	fmt.Println("You are an adult")
	// } else {
	// 	fmt.Println("You are a minor")
	// }
	
	// // if else if example
	// age := 15

	// if age > 18 {
	// 	fmt.Println("You are an adult")
	// } else if age > 12 {
	// 	fmt.Println("You are an teenager")
	// } else {
	// 	fmt.Println("You are a child")
	// }


	// // Logical Operators
	// var role = "admin"
	// var hasPermission = true 

	// if role == "admin" || hasPermission {
	// 	fmt.Println("Access Granted")
	// }

	if age := 13; age > 18 {
		fmt.Println("You are an adult")
	} else if age > 12 {
		fmt.Println("You are an teenager")
	} else {
		fmt.Println("You are minor")
	}


	// go does not have any ternary operator you need use if else  
}