package fizzbuzz

import (
	"fmt"
	"testing"
)

func TestFizzbuzz(t *testing.T) {
	tt := []struct {
		in  int
		out string
	}{
		{1, "1"},
		{2, "2"},
		{3, "Fizz"},
		{4, "4"},
		{5, "Buzz"},
		{6, "Fizz"},
		{7, "7"},
		{8, "8"},
		{9, "Fizz"},
		{10, "Buzz"},
		{11, "11"},
		{12, "Fizz"},
		{13, "13"},
		{14, "14"},
		{15, "FizzBuzz"},
	}
	for _, tst := range tt {
		fb := Fizzbuzz(tst.in)
		if fb != tst.out {
			t.Errorf("fizzbuzz(%v): got %v expected %v\n", tst.in, fb, tst.out)
		}
	}

	for i := 0; i < 100; i++ {
		fb := Fizzbuzz(i)
		if fb != "Fizz" && fb != "Buzz" && fb != "FizzBuzz" && fb != fmt.Sprintf("%d", i) {
			t.Errorf("fizzbuzz(%v): got %v expected Fizz, Buzz, FizzBuzz or %s\n", i, fb, fmt.Sprintf("%d", i))
		}
	}
}
