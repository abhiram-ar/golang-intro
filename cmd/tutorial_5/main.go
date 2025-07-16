package main

import (
	"fmt"
	"strings"
)

func main() {
	myString := "abhéram" // é is a non-acsii characeter
	indexed := myString[1]
	fmt.Printf("%v %T \n", indexed, indexed)

	for i, v := range myString {
		fmt.Printf("idx: %v   value: %v \n", i, v)
	}

	myrune := 'a'
	fmt.Printf("%v %T \n", myrune, myrune)

	strSlice := []string{"s", "t", "r", "i", "n", "g"}
	catStr := ""
	var stringBuilder strings.Builder
	for i := range strSlice {
		// this is inefficient as string is immutable in go and  each operation creates a new string
		catStr += strSlice[i]

		// efficient approch
		// an arary is used under the hood to appemd the string int values
		// string is only when we invoke .string() method on the Builder
		stringBuilder.WriteString(strSlice[i])
	}
	catStr2 := stringBuilder.String()
	fmt.Println(catStr, catStr2)

	// catStr[1] = "a" // strings are immutable
}
