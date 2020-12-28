package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func mix(num int, flag bool, strings ...string) {
	fmt.Println(num, flag, strings)
}

func maxinum(numbers ...float64) float64  {
	max := math.Inf(-1)
	for _, number := range numbers {
		if number > max {
			max = number
		}
	}
	return max
}

func main() {
	arguments := os.Args[1:]
	var sum float64 = 0
	for _, argument := range arguments {
		number, err := strconv.ParseFloat(argument, 64)
		if err != nil {
			log.Fatal(err)
		}
		sum += number
	}
	sampleCount := float64(len(arguments))
	fmt.Printf("Average: %0.2f\n", sum / sampleCount)

	mix(1, true, "a", "b")
	mix(2, false, "a", "b", "c", "d")

	fmt.Println(maxinum(71.8, 56.2, 89.5))
}
