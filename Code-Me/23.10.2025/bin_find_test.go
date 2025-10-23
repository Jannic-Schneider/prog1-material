package searching

import "fmt"

func ExampleBinFind() {
	l1 := []int{1,6,8,10,25,42,125,277}

	fmt.Println(BinFind(l1, 3))
	fmt.Println(BinFind(l1, 8))

	//Output:
	//-1
	//2
}
