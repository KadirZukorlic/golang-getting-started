package main

import "fmt"

type shape interface {
	getArea() float64
}
type square struct {
	sideLength float64
}
type triangle struct {
	height float64
	base   float64
}

func main() {

	tr := triangle{base: 5, height: 7}
	sq := square{sideLength: 4}

	printArea(tr)
	printArea(sq)

}

func printArea(sh shape) {
	fmt.Println(sh.getArea())
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}
