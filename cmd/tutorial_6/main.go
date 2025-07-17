package main

import (
	"fmt"
)

type gasEngine struct {
	mpg     uint8
	galllon uint8
	owner
	int
}

func (e gasEngine) milesLeft() uint8 {
	return e.galllon * e.mpg // method have acess to the struct itself
}

type electicEngine struct {
	mpkwh uint8
	kwh   uint8
	owner
	int
}

func (e electicEngine) milesLeft() uint8 {
	return e.kwh * e.mpkwh
}

type owner struct {
	name string
}

type engine interface {
	milesLeft() uint8
}

func canIMakeIt(e engine, miles uint8) {
	if e.milesLeft() < miles {
		fmt.Println("cannot make it")
	} else {
		fmt.Println("you can make it")
	}
}

func main() {
	var myEngine gasEngine
	fmt.Println(myEngine) // default values

	var myEngine2 gasEngine = gasEngine{mpg: 10, galllon: 2}
	fmt.Println(myEngine2)

	var myEngine3 gasEngine = gasEngine{10, 20, owner{"alex"}, 100} // fields can be ommited
	myEngine3.galllon = 24
	myEngine3.name = "someone"
	fmt.Println(myEngine3)

	// amoymous struct
	myengine4 := struct {
		mpg     uint8
		galllon uint8
	}{1, 10}
	fmt.Println(myengine4)

	fmt.Printf("Miled left for myEngine3 : %v \n", myEngine3.milesLeft())

	myEngine5 := electicEngine{100, 10, owner{"abhira"}, 100}
	canIMakeIt(myEngine2, 128)
	canIMakeIt(myEngine5, 128)
}
