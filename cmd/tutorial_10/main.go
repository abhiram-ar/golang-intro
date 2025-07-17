package main

import "fmt"

func main() {
	intSlice := []int{1, 2, 3, 4, 5}
	res := sumSlice[int](intSlice)
	fmt.Println(res)

	fmt.Println("is slice isEmpty", isEmpty([]string{}))
}

func sumSlice[T int | float32 | float64](nums []T) T {
	var sum T
	for _, val := range nums {
		sum += val
	}
	return sum
}

func isEmpty[T any](slice []T) bool {
	return len(slice) == 0
}
