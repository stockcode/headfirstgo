package main

import "fmt"

func main()  {
	var counters [3]int
	counters[0]++
	counters[0]++
	counters[2]++
	fmt.Println(counters[0], counters[1], counters[2])

	primes := [5]int{2,3,5,7,11}
	fmt.Println(primes)
	fmt.Printf("%#v\n", primes)

	for index, prime := range primes {
		fmt.Println(index, prime)
	}
}
