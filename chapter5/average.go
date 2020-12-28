package main

import (
	"fmt"
	"headfirstgo/datafile"
	"log"
)

func main() {
	numbers, err := datafile.GetFloats("data.txt")
	if err != nil {
		log.Fatal(err)
	}

	sum := 0.0
	for _, number := range numbers{
		sum += number
	}
	sampleCount := float64(len(numbers))

	fmt.Printf("Average: %0.2f\n", sum /sampleCount)
}
