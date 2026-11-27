package main

import (
	"fmt"
)

/*
	range is a keyword ( itteration construct )
	range we use for itterating datastructure
	! range returns that allows you to iterate over elements of data structures.
	!IMP: with range you get two values index and value for array and slices and
	key and value in map
*/ 
func main(){
	// nums := []int {1,2,6,7,8}

	/* how we can itterate in slice ? */
	// for i:=0; i < len(nums); i++ {
	// 	fmt.Println(nums[i])
	// }

	/* Itterate using range */
	// for i, n := range nums {
	// 	fmt.Println(i, n)
	// } 

	/* range in maps */
	// m:= map[string]string{"name":"Arnab", "designation":"SDE"}

	// for k, v := range m {
	// 	fmt.Println(k, v)
	// }

	for i, c := range "golang" {
		fmt.Println(i, string(c))
	}

}