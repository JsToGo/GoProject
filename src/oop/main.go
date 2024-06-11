package main

import "fmt"

// go支持只提供类型而不写字段名的方式，也就是匿名字段，也称为嵌入字段

type Person struct {
	name string
	sex  string
	age  int
}

type Student struct {
	Person
	id   int
	addr string
}

func main() {
	s1 := Student{Person{"5lmh", "man", 20}, 1, "bj"}
	fmt.Println(s1)

	s2 := Student{Person: Person{"5lmh", "man", 20}}
	fmt.Println(s2)

	s3 := Student{Person: Person{name: "5lmh"}, id: 3}
	fmt.Println(s3)

	s4 := Student{Person: Person{name: "r", sex: "man", age: 23}, id: 8, addr: "wh"}
	fmt.Println(s4)
}
