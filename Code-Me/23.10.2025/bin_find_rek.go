package searching

//Liefert die position x in der Liste l,
//oder liefert -1 falls x nicht in l vorkommt.

func BinFindRek(l []int, x int) int {
	if len(l) == 0 {
		return -1
	}
	mitte := len(l) / 2
	if x == l[mitte] {
		return mitte
	}
	if x < l[mitte] {
		//Suche im Linken Teil weiter.
		return BinFindRek(l[:mitte], x)
	}
	if x > l[mitte] {
		//Suche im rechten Teil weiter.
		result := BinFindRek(l[mitte+1:], x)
		if result == -1 {
			return -1 
		}
		return result + mitte +1
	}
	return -1 
}
