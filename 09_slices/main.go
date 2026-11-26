package main

import (
	"fmt"
)

// slice -> dyanmic array
// most used construct in go
// + usefull methods
func main()	{
	// var nums []int 

	// fmt.Println(nums)

	// var nums = make([]int, 2, 5) 
	// // capacity -> maximum numbers of elements can be fit 
	// fmt.Println(nums)

	// // maximum number of elements can be fit 
	// fmt.Println(cap(nums)) // 5

	// fmt.Println(len(nums)) // 2

	// nums = append(nums, 1) 
	// nums = append(nums, 2) 
	// nums = append(nums, 3) 
	// nums = append(nums, 4) 
	// fmt.Println(nums)
	// fmt.Println("length is ", len(nums)) // 6
	// fmt.Println("capacity is ",cap(nums)) // 10 when the capacity exceeds it multiplies the capacity by capacity*2

	// // another declaration of slice
	// nums := []int {}

	// nums = append(nums, 1)
	// nums = append(nums, 2)
	// nums = append(nums, 3)
	// nums = append(nums, 4)
	// fmt.Println(nums)

	/** copy method in slices **/

	// var nums = []int {1,2,3,4,5}
	// nums2 := make([]int, len(nums))

	// fmt.Println(nums, nums2)

	// copy(nums2, nums) // copy the elements of nums to nums2

	// fmt.Println(nums, nums2)

	/** slice operator **/
	// var nums = []int {1,2,3,4,5}
	// fmt.Println(nums[0:1])
	// fmt.Println(nums[0:]) // omiting the end range
	// fmt.Println(nums[:3]) // omitting the start range 


	// slice 

	// var nums1 = []int{1, 2, 3}
	// var nums2 = []int{1, 2, 3}

	// fmt.Println(slices.Equal(nums1, nums2)) // true
	
	// 2D slice 

	var nums = [][]int{{1, 2, 3}, {1, 2, 3}}
	fmt.Println(nums)
	

}