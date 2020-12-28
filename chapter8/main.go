package main

import (
	"fmt"
	"headfirstgo/magazine"
)

type part struct {
	description string
	count int
}

type car struct {
	name string
	topSpeed float64
}

func main() {
	var myStruct struct {
		number float64
		word string
		toggle bool
	}

	myStruct.number = 3.14
	myStruct.word = "pie"
	myStruct.toggle = true

	fmt.Println(myStruct.number)
	fmt.Println(myStruct.word)
	fmt.Println(myStruct.toggle)

	var porsche car
	porsche.name = "Porsche 911 R"
	porsche.topSpeed = 323
	fmt.Println(porsche)

	var bolts part
	bolts.description = "Hex bolts"
	bolts.count = 24
	fmt.Println(bolts)


	subscriber := magazine.Subscriber{Name:        "Aman Singh"	}
	subscriber.City = "omaha"
	subscriber.PostalCode = "12345"
	fmt.Println(subscriber)
}
