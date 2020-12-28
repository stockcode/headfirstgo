package main

import (
	"fmt"
)

func Socailize() error {
	defer fmt.Println("Goodbye!")
	fmt.Println("Hello!")
	return fmt.Errorf("I don't want to talk.")
	fmt.Println("Nice weather, eh?")
	return nil
}

func count(start int, end int)  {
	fmt.Println(start)
	if start < end {
		count(start + 1, end)
	}
}

func calmDown() {
	fmt.Println(recover())
}

func freakOut()  {
	defer calmDown()
	panic("oh no")
}

func main()  {
	//err := Socailize()
	//if err != nil {
	//	log.Fatal(err)
	//}

	//count(1, 3)

	//one()

	freakOut()
	fmt.Println("Exiting normally")
}

func one() {
	two()
}

func two() {
	three()
}

func three()  {
	panic("This call stack's too deep for me!")
}