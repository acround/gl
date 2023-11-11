package record

import "fmt"

type User struct {
	username string
	age      int64
	password string
}

func (user User) isOldEnough() bool {
	return user.age >= 18
}

func run() {
	john := User{username: "John", age: 17, password: "qwerty"}
	mary := User{username: "Mary", age: 21, password: "sdfgh"}
	out := []bool{john.isOldEnough(), mary.isOldEnough()}
	fmt.Println(out)
}
