package main

import "fmt"

/* with variadic functions we can pass unknown numbers of arguments */

func sum(nums ...int) int {
	sum:= 0

	for i:= range nums {
		sum = sum + nums[i]
	}

	return sum
}

func main() {

	// with simple call
	val := sum(1,2,3,4,5,6)
	fmt.Println(val)

	// with spreading 
	nums := []int{1,3,4}
	sum2 := sum(nums...)
	fmt.Println(sum2)

}