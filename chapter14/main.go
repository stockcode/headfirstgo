package main

import (
	"fmt"
	"headfirstgo/datafile"
)

func main() {
	phrases := []string{"my parents", "a rodeo clown"}
	fmt.Println("A photo of", datafile.JoinWithCommas(phrases))
	phrases = []string{"my parents", "a rodeo clown", "a prize bull"}
	fmt.Println("A photo of", datafile.JoinWithCommas(phrases))

	phrases = []string{"my parents"}
	fmt.Println("A photo of", datafile.JoinWithCommas(phrases))
}
