package main

import (
	"fmt"

	"github.com/kaizen63/go-fizzbuzz/fizzbuzz"
)

func main() {
	for i := 1; i <= 100; i++ {
		fmt.Println(fizzbuzz.Fizzbuzz(i))
	}
}
