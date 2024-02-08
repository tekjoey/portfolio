package main

import (
	"fmt"
	"math"
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

// By considering the terms in the Fibonacci sequence whose values do not exceed four million, find the sum of the even-valued terms.
func problemTwo() int {
	fibs := []int{1, 1}
	last := 1
	sum := 0
	for last < 4000000 {
		next := fibs[len(fibs)-1] + fibs[len(fibs)-2]
		fibs[0] = last
		fibs[1] = next
		if next%2 == 0 {
			sum += next
		}
		last = next
	}
	return sum
}

// What is the largest prime factor of the number 600851475143?
func problemThree() int {
	num := 600851475143
	divisor := 2

	for float64(num) > math.Sqrt(float64(num)) {
		//fmt.Println("The current num is:", num)
		//fmt.Println("The divisor is:", divisor)
		if num%divisor == 0 {
			//fmt.Printf("%v / %v == 0\n", num, divisor)
			num /= divisor
		}
		divisor++

	}
	//fmt.Println(divisor)
	return divisor - 1
}

func main() {
	fmt.Println("The answer to problem 1 is:", problemOne())
	fmt.Println("The answer to problem 2 is:", problemTwo())
	fmt.Println("The answer to problem 3 is:", problemThree())
}
