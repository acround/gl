package routine

import (
	"fmt"
	"time"
)

type Lesson struct {
}

func (l Lesson) Run() {
	go l.fact()
	go l.exponenta()
	time.Sleep(2 * time.Second)
}
func (l Lesson) exponenta() {
	val := 1
	for i := 1; i <= 10; i++ {
		val = val * 2
		fmt.Println("Exp(" + fmt.Sprint(i) + ")=" + fmt.Sprint(val))
		time.Sleep(2 * time.Millisecond)
	}
}
func (l Lesson) fact() {
	val := 1
	for i := 1; i <= 10; i++ {
		val = val * i
		fmt.Println("Fact(" + fmt.Sprint(i) + ")=" + fmt.Sprint(val))
		time.Sleep(2 * time.Millisecond)
	}
}

func New() Lesson {
	return *new(Lesson)
}
