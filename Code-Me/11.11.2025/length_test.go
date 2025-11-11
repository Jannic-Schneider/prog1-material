package length

func ExampleLength_conversions() {
	a := FromCentimeters(5)
	b := FromKilometers(5)

	fmt.Println(a.Centimeters)
	fmt.Println(b.Centimeters)

	// Output:
	//5
	//500000
}
