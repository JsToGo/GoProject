package main

import (
	"fmt"
)

// var arr0 [5]int = [5]int{1, 2, 3}
// var arr1 = [5]int{1, 2, 3, 4, 5}
// var arr2 = [...]int{1, 2, 3, 4, 5, 6}
// var strs = [5]string{3: "hello world", 4: "tom"}

func main() {
	type Person struct {
		name string
		age  uint8
		pro  string
	}

	persons := [3]Person{
		{"zhangsan", 20, "driver"},
		{"lisi", 20, "business"},
	}

	for i := 0; i < len(persons); i++ {
		if i == 2 {
			persons[i] = Person{"wuhu", 20, "thief"}
		}
	}

	for n, v := range persons {
		fmt.Println(n, v)
	}
}
