package main

import "fmt"

func main() {
	var num int = 0
	ptr := &num
	*ptr = 3
	fmt.Println(num)
}
