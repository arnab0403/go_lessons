package main

import (
	"fmt"
	"time"
)

func main() {
	// simple switch case
	i := 5

	switch i {
	case 1:
		fmt.Println("One")
	case 2: 
	    fmt.Println("Two")
	case 3: 
	    fmt.Println("Three")
	default:
		fmt.Println("Please enter a valid number")
	}

	/**  multiple condition switch case **/ 
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("This is a sunday")
	
	default: 
		fmt.Println("It's a working day")
	}

	// type switch 
	whoAmI := func(i interface{}){
		switch i.(type) {
		case int:
			fmt.Println("It's an integer")
		case string:
			fmt.Println("It's a string")
		case bool:
			fmt.Println("It's a boolean")
		default:
			fmt.Println("Others")
		}
	}

	whoAmI(5.5)
	whoAmI("hello")
	whoAmI(true)
}
