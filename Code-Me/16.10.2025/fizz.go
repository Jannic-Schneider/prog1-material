package fizzbuzz

import "fmt"

// Fizz gibt die Zahlen von 1 bis 30 aus und ersetzt dabei
// jede durch 3 teilbare Zahl durch fizz und
// jede durch 5 teilbare zahl durch buzz
// Bei zahlen die durch 3 und 5 Teilbar sind wird Fizzbus ausgegeben

func Fizz() {
	for i := 1; i <= 30; i++ {

		//wenn i durch 3 und 5 teilbar ist, gib "fizzbus" aus
		if i%5 == 0 && i%3 == 0 {
			fmt.Println("FizzBuzz")
		} else
		//wenn i durch 3 teilbar ist, gib "fizz" aus
		if i%3 == 0 {
			fmt.Println("Fizz")
		} else
		//wenn i durch 5 teilbar ist, gib "buzz" aus
		if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}

	}

}
