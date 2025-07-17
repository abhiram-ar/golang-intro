package main

import "fmt"

func main() {
	var p *int32 = new(int32)
	var i int32

	fmt.Printf("value of p pointing to %v \n", *p) // pointer dereferencing
	fmt.Printf("value of i %v \n", i)

	*p = 10
	fmt.Printf("value of p pointing to %v \n", *p) // pointer dereferencing

	p = &i                                         // points to mem location of i
	fmt.Printf("value of p pointing to %v \n", *p) // pointer dereferencing

	*p = 100
	fmt.Printf("value of p pointing to %v \n", *p) // pointer dereferencing
	fmt.Printf("value of i %v \n", i)

	var arr [4]int32 = [4]int32{1, 2, 3, 4}
	fmt.Printf("memory adress of input array %p \n", &arr)
	reult := square(&arr)
	fmt.Printf("result arr: %v", reult)
}

func square(nums *[4]int32) [4]int32 {
	fmt.Printf("address of function array %p \n", nums)
	for i := range nums {
		nums[i] = nums[i] * nums[i]
	}

	return *nums
}
