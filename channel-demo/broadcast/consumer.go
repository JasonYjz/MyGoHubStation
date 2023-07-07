package main

import (
	"context"
	"fmt"
)

type Traffic interface {
	way()
}

type Train struct {
	name     string
	personCh chan string
	ctx      context.Context
}

func NewTrain(name string, personCh chan string, ctx context.Context) *Train {
	return &Train{name: name, personCh: personCh, ctx: ctx}
}

func (t Train) way() {
	for {
		select {
		case per := <-t.personCh:
			fmt.Printf("%s use %s way\n", per, t.name)
		case <-t.ctx.Done():
			fmt.Printf("%s way Done\n", t.name)
			return
		}
	}
}

type Car struct {
	name     string
	personCh chan string
	ctx      context.Context
}

func NewCar(name string, personCh chan string, ctx context.Context) *Car {
	return &Car{name: name, personCh: personCh, ctx: ctx}
}

func (c Car) way() {
	for {
		select {
		case per := <-c.personCh:
			fmt.Printf("%s use %s way\n", per, c.name)
		case <-c.ctx.Done():
			fmt.Printf("%s way Done\n", c.name)
			return
		}
	}
}

type Plane struct {
	name     string
	personCh chan string
	ctx      context.Context
}

func NewPlane(name string, personCh chan string, ctx context.Context) *Plane {
	return &Plane{name: name, personCh: personCh, ctx: ctx}
}

func (p Plane) way() {
	for {
		select {
		case per := <-p.personCh:
			fmt.Printf("%s use %s way\n", per, p.name)
		case <-p.ctx.Done():
			fmt.Printf("%s way Done\n", p.name)
			return
		}
	}
}
