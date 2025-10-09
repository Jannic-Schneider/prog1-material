package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Du hast drei versuche die Richtige Zahl zwischen 0 und 100 zu erraten!")

	//my_number... Gibt eine zufällige zahl zwischen 0 und n (hier n + 1 )
	my_number := rand.Intn(100) + 1

	//for i ist eine Schleife, Zählt 1,2,3 eingaben und bricht dann ab.
	for i := 0; i < 3; i++ {
		guess := ReadNumber()

		//if NumberGood zieht sich die Zahl aus der funktion NumberGood(zufalls generierte zahl)
		//und vergleicht sie mit der eingegebene Zahl

		if NumberGood(guess, my_number) {
			fmt.Println("Das war richtig!")
			return
		}
		fmt.Printf("%d war nicht korrekt!\n", guess)
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

// NumberGood prüft, ob x gleich einer Zufälligen Zahl zwischen 1 und 100 ist.
func NumberGood(x int, n int) bool {

	return x == n

}
