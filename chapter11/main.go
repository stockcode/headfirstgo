package main

import "fmt"

type Whistle string

func (w Whistle) MakeSound()  {
	fmt.Println("Tweet!")
}

type Horn string

func (h Horn) MakeSound()  {
	fmt.Println("Honk!")
}

type NoiseMaker interface {
	MakeSound()
}

func play(n NoiseMaker)  {
	n.MakeSound()
}

type ComedyError string

func (c ComedyError) Error() string {
	return string(c)
}

func main() {
	var toy NoiseMaker
	toy = Whistle("Toyco Canary")
	toy.MakeSound()

	toy = Horn("Toyco Blaster")
	toy.MakeSound()

	play(Whistle("Toyco Canary"))
	play(Horn("Toyco Blaster"))

	var err error
	err = ComedyError("What's a programmer's favoriate beer?")
	fmt.Println(err)
}
