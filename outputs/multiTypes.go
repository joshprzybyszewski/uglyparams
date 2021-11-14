package example

import "fmt"

func TwoTypes(
	a, b string,
	x, y int,
) {
	println(a)
	println(b)
	fmt.Printf("%d %d\n", x, y)
}

func ThreeTypes(
	a, b string,
	x, y int,
	u, v float64,
) {
	println(a)
	println(b)
	fmt.Printf("%d %d\n", x, y)
	fmt.Printf("%v %v\n", u, v)
}

func FourTypes(
	a, b string,
	x, y int,
	u, v float64,
	b1, b2 bool,
) {
	println(a)
	println(b)
	fmt.Printf("%d %d\n", x, y)
	fmt.Printf("%v %v\n", u, v)
	fmt.Printf("%v %v\n", b1, b2)
}
