package length

type Length interface {}

func FromCentimeters(c int) Length {

	return Length(c)
}
func FromMeters(m int) Length {

	return Length(m *100)
}
func FromKilometers(k int) Length {

	return Length(k * 1000)
}