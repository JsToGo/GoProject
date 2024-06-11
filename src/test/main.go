package main

import (
	"fmt"
)

type Sayer interface {
	say()
}

type Mover interface {
	move()
}

type Animal interface {
	Sayer
	Mover
}

type cat struct {
	name string
}

func (c cat) say() {
	fmt.Println("喵喵喵")
}

func (c cat) move() {
	fmt.Println("猫会动")
}

func main() {
	var x Animal = cat{name: "hh"}
	x.say()
	x.move()
}
