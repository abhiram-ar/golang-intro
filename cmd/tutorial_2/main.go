package main

import "fmt"

func main() {
	var a int8 = 10
	fmt.Println(a)

	// var b int8 = 1001313
	// fmt.Println(b) // this throw error as int8 cannnot hold such a large number

	var str string = "abhiram"
	fmt.Println(str)

	str2 := "short declaration operator, intertype from right hand side"
	fmt.Println(str2)

	r := '😑' // go uses utf-8 for string encoding, to handle unicode wich can have more than one byte per character we use runes,
	fmt.Println(r)
}
