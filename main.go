package main

import (
	"fmt"
)

func Fizzbuzz(i int) string {
	var result string
	if i%15 == 0 {
		result = "FizzBuzz"
	} else if i%5 == 0 {
		result = "Buzz"
	} else if i%3 == 0 {
		result = "Fizz"
	} else {
		result = fmt.Sprint(i)
	}
	return result
}
func main() {
	for i := 1; i <= 100; i++ {
		fmt.Println(Fizzbuzz(i))
	}
}
