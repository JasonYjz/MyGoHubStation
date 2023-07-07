package main

import "fmt"

type Breath interface {
	Way()
}

type Fish struct {
	name string
}

func (f *Fish) Way() {
	fmt.Printf("%s name\n", f.name)
}

type Human struct {
	age string
}

func (h *Human) Way() {
	fmt.Printf("%s human\n", h.age)
}
