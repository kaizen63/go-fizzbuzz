package main

import (
	"fmt"
	"./fizzbuzz"
)

func main() {
	for i := 1; i <= 100; i++ {
		fmt.Println(fizzbuzz.Fizzbuzz(i))
	}
}
