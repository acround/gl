package record

import "fmt"

type User struct {
	username string
	age      int64
	password string
}

type Lesson struct {
}

func (user User) getUserName() string {
	return user.username
}

func (user User) getAge() int64 {
	return user.age
}

func (user User) isOldEnough() bool {
	return user.age >= 18
}

func (l Lesson) Run() {
	company := []User{
		User{username: "John", age: 17, password: "qwerty"},
		User{username: "Mary", age: 21, password: "sdfgh"},
	}
	for _, user := range company {
		fmt.Printf("%s, %d years old -> ", user.getUserName(), user.getAge())
		if user.isOldEnough() {
			fmt.Println("allowed")
		} else {
			fmt.Println("denied")
		}

	}
}

func New() Lesson {
	return *new(Lesson)
}
