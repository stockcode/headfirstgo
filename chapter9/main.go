package main

import "fmt"

type Liters float64
type Gallons float64
type Milliliters float64

func (l Liters) toGallons() Gallons  {
	return Gallons(l * 0.264)
}

func (m Milliliters) toGallons() Gallons  {
	return Gallons(m * 0.000264)
}

type MyType string

type Number int

func (m MyType) sayHi()  {
	fmt.Println("Hi from", m)
}

func (n Number) Add(otherNumber int) {
	fmt.Println(n, "plus", otherNumber, "is", int(n) + otherNumber)
}

func (n Number) Substract(otherNumber int) {
	fmt.Println(n, "minus", otherNumber, "is", int(n) - otherNumber)
}

func (n *Number) Double()  {
	*n *= 2
}

func main()  {
	var carFuel Gallons
	var busFuel Liters

	carFuel = Gallons(Liters(40.0) * 0.264)
	busFuel = Liters(240.0)

	fmt.Println(carFuel, busFuel)

	value := MyType("a MyType value")
	value.sayHi()

	ten := Number(10)
	ten.Add(4)
	ten.Substract(5)

	four := Number(4)
	four.Add(3)
	four.Substract(2)

	number := Number(4)
	fmt.Println("Original value of number:", number)
	number.Double()
	fmt.Println("number after calling Double:", number)

	soda := Liters(2)
	fmt.Printf("%0.3f liters equals %0.3f gallons\n", soda, soda.toGallons())
	water := Milliliters(500)
	fmt.Printf("%0.3f milliliters equals %0.3f gallons\n", water, water.toGallons())
}
