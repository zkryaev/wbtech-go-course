package main

import "fmt"

type Human struct {
	gender string
	age    uint
	name   string
}

func (h Human) GetName() string {
	return h.name
}

func (h Human) GetAge() uint {
	return h.age
}

func (h Human) GetGender() string {
	return h.gender
}

type Action struct {
	Human
}

func main() {
	action := &Action{
		Human{"megamen", 45, "igor"},
	}
	fmt.Println(action.GetName(), action.GetGender(), action.GetAge())
}
