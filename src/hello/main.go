package main

import (
	"fmt"
	"math"
)

const (
	n1 int    = iota
	n2 string = "123"
	n3 int    = iota
)

var flag bool

func foo(int, string) (int, string) {
	return 10, "Q1m1"
}

func changeString() {
	s1 := "hello"
	// 强制类型转换
	byteS1 := []byte(s1)
	byteS1[0] = 'H'
	fmt.Println(byteS1)
	fmt.Println(string(byteS1))

	s2 := "博客"
	runeS2 := []rune(s2)
	runeS2[0] = '狗'
	fmt.Println(runeS2)
	fmt.Println(string(runeS2))
}

func sqrtDemo() {
	var a, b = 3, 4
	// math.Sqrt()接收的参数是float64类型，需要强制转换
	c := int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

func main() {
	sqrtDemo()
	fmt.Println("-------------------------")
	changeString()
	fmt.Println("-------------------------")
	x, _ := foo(1, "123")
	_, y := foo(1, "123")

	s1 := `
第一行
第二行
第三行
    `
	fmt.Println("x =", x)
	fmt.Println("y =", y)
	fmt.Println("n1 =", n1)
	fmt.Println("n2 =", n2)
	fmt.Println("n3 =", n3)
	fmt.Println("flag =", flag)
	fmt.Println("str := \"c:\\pprof\\main.exe\"")
	fmt.Println(s1)

}
