package main

import "fmt"

func main() {
	array3 := [5]string{"a", "b", "c", "d", "e"}
	slice3 := array3[0:3]
	slice4 := array3[2:5]
	fmt.Println(slice3, slice4)

	var slice []string
	if len(slice) == 0 {
		slice = append(slice, "first item")
	}
	fmt.Printf("%#v\n", slice)
}
