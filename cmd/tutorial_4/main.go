package main

import "fmt"

func main() {
	var intArr [3]int32
	intArr[0] = 123
	fmt.Println(intArr[0])
	fmt.Println(intArr[1:3])                        // eleements from index 1 to 3 (excluded)
	fmt.Println(&intArr[0], &intArr[1], &intArr[2]) // memory location of each array elements (contigious 4 bytes)

	var intArr2 [3]int32 = [3]int32{1, 2, 3} // inmmedidate array inittialziation
	fmt.Println(intArr2)

	intArr3 := [...]int32{1, 2, 3} // compilter infers the array size
	fmt.Println(intArr3)

	// ----- slices -------------
	var intSlice []int32 = []int32{4, 5, 6}
	fmt.Printf("%v has length %v and capacity %v \n", intSlice, len(intSlice), cap(intSlice))

	intSlice = append(intSlice, 5)
	fmt.Printf("%v has length %v and capacity %v \n", intSlice, len(intSlice), cap(intSlice))

	// tough the new capacity of the slice is now 6 if we try to access elements larger than lenght,
	// we will get an index out of rannge error
	// intSlice[5] = 10

	intSlice2 := []int32{10, 11}
	intSlice = append(intSlice, intSlice2...) // spread operator
	fmt.Println(intSlice)

	// another way to make slice - using make function
	intSlice3 := make([]int32, 3, 10) // makes a slice of type `int32`, length - 3 and capacity 10 (default to length)
	fmt.Println(intSlice3)

	// ------------map-------------
	myMap := make(map[string]uint)
	fmt.Println(myMap)

	myMap2 := map[string]uint{"Abhiram": 24, "Ishan": 19, "suku": 25}
	fmt.Println(myMap2)
	fmt.Println(myMap2["Abhiram"])
	fmt.Println(myMap2["aaroo"])

	age, ok := myMap2["aaroo"]
	if ok {
		fmt.Printf("elemenet found %v \n", age)
	} else {
		fmt.Printf("Item not found \n")
	}

	delete(myMap2, "Abhiram") // delete by refrence, so no return value
	fmt.Println(myMap2)

	// ---------iteration -----------
	// for array, slice, map
	for name, age := range myMap2 {
		fmt.Println(name, age)
	}
}
