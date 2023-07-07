package main

type My struct {
	id int
	b  Breath
}

func NewMy(id int) *My {
	return &My{
		id: id,
		b:  &Human{age: "18"},
	}
}

func main() {
	my := NewMy(1)
	my.b.Way()

	my.b = &Fish{name: "bigfish"}
	my.b.Way()
}
