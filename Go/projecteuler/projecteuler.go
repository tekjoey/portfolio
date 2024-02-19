package main

import (
	"fmt"
	"math"
	"strconv"
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

// Find the largest palindrome made from the product of two 3-digit numbers.
func problemFour() int {

	// A helper function to [det]ermine if a number is a [pal]endrome
	detPal := func(inte int) bool {
		str0 := strconv.Itoa(inte)

		// Refactor to optimize
		// for i := 0; i < len(str0); i++ {
		// 	if st0[i]!= str0[len(str0)]
		// }

		halfLength := len(str0) / 2
		str1 := str0[0:halfLength]

		// If length is odd, dont need to include the middle digit.
		var str2 string
		if len(str0)%2 == 0 {
			str2 = str0[halfLength:]
		} else {
			str2 = str0[halfLength+1:]
		}

		// reverse the strings
		str2RevRune := []rune(str2)
		for i, j := 0, len(str2RevRune)-1; i < j; i, j = i+1, j-1 {
			str2RevRune[i], str2RevRune[j] = str2RevRune[j], str2RevRune[i]
		}
		str3 := string(str2RevRune)

		if str1 == str3 {
			return true
		} else {
			return false
		}
	}

	var highest int
	for i := 100; i <= 999; i++ {
		for j := 100; j <= 999; j++ {
			product := i * j
			if detPal(product) {
				if product > highest {
					highest = product
				}

			}
		}
	}
	return highest
}

func main() {
	fmt.Println("The answer to problem 1 is:", problemOne())
	fmt.Println("The answer to problem 2 is:", problemTwo())
	fmt.Println("The answer to problem 3 is:", problemThree())
	fmt.Println("The answer to problem 4 is:", problemFour())
}
