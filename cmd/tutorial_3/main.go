package main

import (
	"errors"
	"fmt"
)

func main() {
	inputStr := "hello world"
	printMe(inputStr)

	q, r, error := intDivison(10, 0)
	if error != nil {
		fmt.Println(error.Error())
	} else if r == 0 {
		fmt.Printf("The quotiant is %v ", q)
	} else {
		fmt.Printf("The quotiant is %v and remainter is %v", q, r)
	}

	// using swithc statement
	switch {
	case error != nil:
		fmt.Println(error.Error())
	case r == 0:
		fmt.Printf("The quotiant is %v ", q)
	default:
		fmt.Printf("The quotiant is %v and remainter is %v", q, r)
	}
}

func printMe(input string) {
	fmt.Println(input)
}

func intDivison(numerator int, denominator int) (int, int, error) {
	var err error // default value of error is nill
	if denominator == 0 {
		err = errors.New("cannot divide by zero")
		return 0, 0, err

	}
	quotiant := numerator / denominator
	remainter := numerator % denominator
	return quotiant, remainter, err
}
