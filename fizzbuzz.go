package fizzbuzz

import (
	"strconv"
)

// var FizzBuzz = fizzbuzz

func FizzBuzz(n int) string {
	// Atoi
	// Itoa
	if n%15 == 0 {
		return "fizzbuzz"
	}
	if n%3 == 0 {
		return "fizz"
	}
	if n%5 == 0 {
		return "buzz"
	}
	return strconv.Itoa(n)

}
