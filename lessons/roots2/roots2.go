package roots2

import (
	"fmt"
	"math"
)

type Lesson struct {
}

func (l Lesson) Run() {
	var a float64
	var b float64
	var c float64
	var det float64
	fmt.Println("Put coefficients")
	a = l.getCoef("A")
	b = l.getCoef("B")
	c = l.getCoef("C")
	det = l.determinant(a, b, c)
	if det > 0 {
		x1 := (-b - math.Sqrt(det)) / (2 * a)
		x2 := (-b + math.Sqrt(det)) / (2 * a)
		fmt.Println("Roots:")
		fmt.Println("x1=" + fmt.Sprint(x1))
		fmt.Println("x2=" + fmt.Sprint(x2))
	} else if det == 0 {
		x := -b / (2 + a)
		fmt.Println("Root:")
		fmt.Println("x=" + fmt.Sprint(x))
	} else {
		fmt.Println("No root")
	}
}

func (l Lesson) getCoef(name string) float64 {
	var tmp float64
	fmt.Println(name + ":")
	fmt.Scan(&tmp)
	return tmp
}

func (l Lesson) determinant(a float64, b float64, c float64) float64 {
	return (b * b) - (4 * a * c)
}

func New() Lesson {
	return *new(Lesson)
}
