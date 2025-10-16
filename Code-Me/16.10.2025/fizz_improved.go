package fizzbuzz

import "fmt"

func FizzImproved(x, y, n int) {
	for i := 1; i <= 30; i++ {

		if i%5 == 0 && i%3 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else {
			fmt.Println(i)

		}

	}

}
