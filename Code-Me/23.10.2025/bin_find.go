package searching

// Liefert die Position x in der Liste l
// oder liefert -1, falls x nicht in l vorkommt.
func BinFind(l []int, x int) int {
	for {
		mitte := len(l) / 2
		//vergleiche x mit dem Element in der Mitte.
		//wenn x == l[mitte], dann fertig
		if x == l[mitte] {
			return mitte
		}
		//Wenn x < l[mitte], dann suche
		//im Linken teil Weiter.
		if x < l[mitte] {
			l = l[0:mitte]
		} else {
			l = l[mitte+1:]
		}
		//Wenn x > l[mitte], dann suche im Rechten teil Weiter.
		return -1
	}
}
