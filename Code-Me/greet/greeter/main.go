package main

import (
	"fmt"
	"prog1-material/Code-Me/greet"
)

func main() {
	fmt.Println(greet.Greet("Kurs"))
	fmt.Println("Alle " + greet.Greet("Kursteilnehmer"))

}
