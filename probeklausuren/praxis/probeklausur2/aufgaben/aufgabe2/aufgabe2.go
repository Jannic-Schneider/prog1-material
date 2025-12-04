package aufgabe2

/* AUFGABENSTELLUNG: Vervollständigen Sie unten stehende Funktion.
 * ERREICHBARE PUNKTE: 10
 */

// FilterDigits erwartet einen String s und liefert einen String zurück,
// der aus s entsteht, indem alle Ziffern entfernt werden.
// Alle anderen Zeichen sollen unverändert bleiben.
func FilterDigits(s string) string {
	result := ""
	digits := "0123456789"
	for _, el := range s {
		a := true
		for _, za := range digits {
			if el == za {
				a = false
			}
		}

		if a {
			result += string(el)
		}
	}
	return result
}
