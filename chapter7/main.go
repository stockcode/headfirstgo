package main

import "fmt"

func main()  {
	ranks := map[string]int{"bronze": 3, "silver": 2, "gold": 1}
	fmt.Println(ranks["gold"])
	fmt.Println(ranks["bronze"])

	counters := map[string]int{"a": 3, "b": 0}
	var value int
	var ok bool
	value, ok = counters["a"]
	fmt.Println(value, ok)

	value, ok = counters["b"]
	fmt.Println(value, ok)

	value, ok = counters["c"]
	fmt.Println(value, ok)
}
