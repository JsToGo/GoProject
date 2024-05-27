package main

import (
	"fmt"
	"math/rand"
	"time"
)

func sumArr(arr []int) int {
	fmt.Println(arr)
	var res int
	for _, value := range arr {
		res += value
	}
	return res
}

func genRandArray() []int {
	rand.Seed(time.Now().Unix())

	var randArr = make([]int, rand.Intn(20))

	for i := 0; i < len(randArr); i++ {
		randArr[i] = rand.Intn(100)
	}

	return randArr
}

func main() {
	fmt.Println(sumArr(genRandArray()))
}
