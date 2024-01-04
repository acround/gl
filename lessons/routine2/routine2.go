package routine2

import "fmt"

type Lesson struct {
}

func (l Lesson) Run() {
	ch := make(chan int)
	go l.do(ch)
	for a := range ch {
		fmt.Println(a)
	}
}
func (l Lesson) do(ch chan int) {
	for i := 1; i <= 10; i++ {
		ch <- i
	}
	close(ch)
}
func New() Lesson {
	return *new(Lesson)
}
