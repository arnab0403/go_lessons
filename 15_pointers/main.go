package main

import "fmt"

// func changeNum(num int) {
// 	num = 5
// 	fmt.Println("In change Num ", num)
// }


// by reference 
func changeNum(num *int) {
	/* 
		dereference
		!IMP: dereferencing a pointer means accessing the actual value stored at the memory address that the pointer holds. 
	*/
	*num = 5
	fmt.Println("In change Num", *num)
}

func main() {
	num := 1
	fmt.Println("Before change num", num)
	changeNum(&num)
	fmt.Println("Memory address", &num)
}