package main

import (
	"fmt"
)

var arr = [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

func genSlice() {
	// 切片操作
	// slice := arr[start:end:max]
	// start 	切片截取数组索引开始 省略时 start = 0
	// end		切片截取数组索引结尾 省略时 end = len(arr) - 1
	// max		控制切片cap长度 cap(slice) = max - low, 省略时不做约束
	slice1 := arr[1:5]
	slice2 := arr[1:]
	slice3 := arr[:5]
	sliceAll := arr[:]
	sliceCap := arr[0 : len(arr)-1 : 9]
	fmt.Println("slice1		", slice1)
	fmt.Println("slice2		", slice2)
	fmt.Println("slice3		", slice3)
	fmt.Println("sliceAll	", sliceAll)
	fmt.Println("sliceCap	", sliceCap)

}

func genSliceByMake() {
	// make(Type, len, cap)
	// 满足 len <= cap
	s1 := make([]int, 2)
	s2 := make([]int, 10, 12)
	fmt.Println("s1", s1)
	fmt.Println("s2", s2)
}

func compareSliceArray() {
	nums := [...]int{0, 1, 2, 3, 4, 5}

	// 数组 --- 只对值进行复制
	copyNums := nums
	copyNums[3] = 666
	fmt.Println(nums)
	fmt.Println(copyNums)

	// 指针 --- 对地址的引用
	ptr := &nums[3]
	*ptr = 666
	fmt.Println(nums)
	fmt.Println(nums)

	// 切片 --- 对地址的引用
	slice := nums[0:2]
	slice[0] += 100
	slice[1] += 100
	fmt.Println(nums)
	fmt.Println(slice)
}

func modifySlice(innerSlice []string) {
	innerSlice[0] = "b"
	// 如果指向底层数组的指针被覆盖或者修改(copy, 重分配, append触发扩容), 此时函数内部对数据的修改将不再影响到外部的切片
	// 此处触发append后, 触发了切片的扩容，新建了一个底层数组
	innerSlice = append(innerSlice, "a")
	innerSlice[1] = "b"
	fmt.Println("innerSlice", innerSlice)
}

func appendNewSlice() {
	slice1 := make([]int, 0, 5)
	fmt.Printf("%p\n", &slice1)

	// append: 向 slice 尾部添加数据, 返回新的 slice 对象, 地址不同
	slice2 := append(slice1, 2)
	fmt.Printf("%p\n", &slice2)

	fmt.Println(slice1, slice2)
}

func sliceOverCap() {
	arr := [...]int{0, 1, 2, 3, 4, 10: 0}
	slice := arr[0:3:3]
	slice = append(slice, 100, 200) // 一次 append 两个值，超出 s.cap 限制
	fmt.Println(slice, arr)         // 重新分配底层数组，与原数组无关。
	fmt.Println(&slice[0], &arr[0]) // 比对底层数组起始指针。
}

func sliceCap() {
	slice := make([]int, 0, 1)
	c := cap(slice)

	for i := 0; i < 50; i++ {
		slice = append(slice, i)
		if n := cap(slice); n > c {
			fmt.Printf("cap: %d -> %d\n", c, n)
			c = n
		}
	}
}

func main() {
	fmt.Println("-------------genSlice-------------------")
	genSlice()
	fmt.Println("-------------genSliceByMake-------------")
	genSliceByMake()
	fmt.Println("-------------compareSliceArray----------")
	compareSliceArray()
	fmt.Println("-------------modifySlice----------------")
	outerSlice := []string{"a", "a"}
	modifySlice(outerSlice)
	fmt.Println("outerSlice", outerSlice)
	fmt.Println("-------------appendNewSlice-------------")
	appendNewSlice()
	fmt.Println("-------------sliceOverCap---------------")
	sliceOverCap()
	fmt.Println("-------------sliceCap-------------------")
	sliceCap()
}
