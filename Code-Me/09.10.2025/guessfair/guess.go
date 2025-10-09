package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Du hast drei versuche die Richtige Zahl zwischen 0 und 100 zu erraten!")

	my_number := rand.Intn(100) + 1

	for i := 0; i < 3; i++ {
		guess := ReadNumber()
		if guess < my_number {
			fmt.Println("Die gesuchte zahl ist größer!")
		}
		if guess > my_number {
			fmt.Println("Die gesuchte zahl ist kleiner!")
		}
		if guess == my_number {
			fmt.Println("Das war richtig!")
			return
		}
	}
	fmt.Println("Game Over!")

}

// ReadNumber liefer uns ein int.
func ReadNumber() int {
	var n int
	fmt.Print("Bitte gib eine Zahl ein: ")
	//fmt.Scan wartet auf die Eingabe des Nutzers
	fmt.Scan(&n)
	//return gibt die eigegebene Zahl an das System zurück
	return n

}
