package main

import (
	"fmt"
)

func ReadNumber(int) {
	fmt.Print("Rate eine Zahl: ")
	var n int
	fmt.Scan(&n)
	return n
}

func GuessingGame() {
	for n := 0; n < 3; n++ {
		guess := ReadNumber()
		if NumberGood(guess) {
			fmt.Println("Richtig geraten! ")
			return
		} else {
			fmt.Println("Leider Falsch!")
		}
	}
	fmt.Println("Zu viele falsche Zahlen! ")
}

func NumberGood(g, e int) bool {
	return g == e
}

func main() {
	GuessingGame()
}
