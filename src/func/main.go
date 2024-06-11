package main

import "fmt"

type Person struct {
	name string
	age  int
}

func (p Person) say() {
	fmt.Println(p)
}

func (p *Person) grow(i int) {
	p.age += i
}

func main() {
	me := Person{"r", 23}
	me.say()
	me.grow(1)
	me.say()
}
