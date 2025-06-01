package main

import (
	"fmt"
	"reflect"
	"slices"
)

// slice is array of dynamic data types
// Slices, are passed by reference to a function, which means that the function can modify the underlying array.
func main() {
	// It's a slice, not an array. If you pass a size, then it's treated as an array.
	//An uninitialized slice is `nil`, which means `{null}`

	var nums []int
	var nums1 = make([]int, 2, 5) // 2 is initial size and 5 is initial capacity both can automatically increase after adding more elements.

	fmt.Println(nums == nil)  //true
	fmt.Println(nums1 == nil) //false

	//Type checking
	// fmt.Println(reflect.TypeOf(nums).Kind()) // slice

	println(len(nums))  //0
	println(len(nums1)) //2
	println(cap(nums1)) //5
	// adding elements
	nums1 = append(nums1, 3, 3, 2)
	nums = append(nums, 5, 6, 7)
	fmt.Println(nums1) // [0 0 3 3 2]
	fmt.Println(nums)  //[5 6 7]
	nums[0] = 100
	fmt.Println(nums) //[100 6 7]

	nums2 := make([]int, len(nums))
	// To copy the contents of one slice to another, you need to ensure that the destination slice has enough capacity to hold the elements || so the nums2 has to be initialize with size
	copy(nums2, nums)
	fmt.Println(nums2) //[100 6 7 8]

	//slice operator
	n := nums1[2:]
	fmt.Println(n)                        //[3 3 2]
	fmt.Println(reflect.TypeOf(n).Kind()) // slice

	//slices package in GO
	//-In Go, the slices package is not a built-in package, but rather a third-party package that provides additional functionality for working with slices.
	nums3 := []int{1, 5, 3, 4, 9, 7}
	slices.Sort(nums3)
	fmt.Println(nums3)

}
