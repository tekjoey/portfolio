package main

import (
	"fmt"
)

// Find the sum of all the multiples of 3 or 5 below 1000
func problemOne() int {
	var sum int
	for i := 0; i < 1000; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}
	return sum
}

func main() {
	fmt.Println("The answer to problem 1 is: ", problemOne())
}
