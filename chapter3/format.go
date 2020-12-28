package main

import (
	"fmt"
)

func paintNeeded(width float64, height float64) (float64, error) {
	if width < 0 {
		return 0, fmt.Errorf("a width of %0.2f is invalid", width)
	}
	if height < 0 {
		return 0, fmt.Errorf("a height of %0.2f is invalid", width)
	}
	area := width * height
	return area / 10.0, nil
}

func negate(myBoolean *bool) {
	*myBoolean = !*myBoolean
}

func main()  {
	fmt.Printf("A float: %f\n", 3.1415)
	fmt.Printf("An integer: %d\n", 15)
	fmt.Printf("A string: %s\n", "hello")
	fmt.Printf("A boolean: %t\n", false)
	fmt.Printf("A float: %v %v %v\n", 1.2, "\t", true)
	fmt.Printf("A float: %#v %#v %#v\n", 1.2, "\t", true)
	fmt.Printf("A float: %T %T %T\n", 1.2, "\t", true)
	fmt.Printf("A float: Percent sign: %%\n")

	fmt.Printf("%12s | %s\n", "Product", "Cost in Cents")
	fmt.Println("-------------------------------")
	fmt.Printf("%12s | %2d\n", "Stamps", 50)
	fmt.Printf("%12s | %2d\n", "Paper Clips", 5)
	fmt.Printf("%12s | %2d\n", "Tape", 99)

	fmt.Printf("%%7.3f: %7.3f\n", 12.3456)
	fmt.Printf("%%7.2f: %7.2f\n", 12.3456)
	fmt.Printf("%%7.1f: %7.1f\n", 12.3456)
	fmt.Printf("%%.1f: %.1f\n", 12.3456)
	fmt.Printf("%%.2f: %.2f\n", 12.3456)

	var amount, total float64
	amount, _ = paintNeeded(4.2, 3.0)
	fmt.Printf("%0.2f liters needed\n", amount)
	total += amount
	amount, _ = paintNeeded(5.2, 3.5)
	fmt.Printf("%0.2f liters needed\n", amount)
	total += amount
	fmt.Printf("Total: %0.2f liters\n", total)

	//err := errors.New("height can't be negative")
	//fmt.Println(err)
	//log.Fatal(err)

	//err := fmt.Errorf("a height of %0.2f is invalid", -2.33333)
	//fmt.Println(err)

	amount, err := paintNeeded(4.2, -3.0)
	fmt.Println(err)
	fmt.Printf("%0.2f liters needed\n", amount)

	amount = 6
	fmt.Println(amount)
	fmt.Println(&amount)

	myInt := 4
	myIntPointer := &myInt
	*myIntPointer = 8
	fmt.Println(myIntPointer)
	fmt.Println(*myIntPointer)

	myFloat := 98.6
	myFloatPointer := &myFloat
	fmt.Println(myIntPointer)
	fmt.Println(*myFloatPointer)

	myBool := true
	myBoolPointer := &myBool
	fmt.Println(myBoolPointer)
	fmt.Println(*myBoolPointer)

	truth := true
	negate(&truth)
	fmt.Println(truth)

	lies := false
	negate(&lies)
	fmt.Println(lies)
}
