package main

import "fmt"

type user struct {
	name string
	age  uint8
}

func (u user) save() {
	fmt.Println("User saved")
}

func (u user) isTeenager() bool {
	return u.age >= 18
}

func (u *user) doingBirthDay() {
	u.age++
}

func main() {
	user1 := user{"Usuário", 23}
	fmt.Println(user1)

	user1.save()
	user1.isTeenager()
	user1.doingBirthDay()
	fmt.Println(user1.age)
}
