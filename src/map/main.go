package main

import "fmt"

//	map[KeyType]ValueType
// 	KeyType:表示键的类型。
// 	ValueType:表示键对应的值的类型。
//	make(map[KeyType]ValueType, cap)

func main() {
	// scoreMap := map[string]int{
	// 	"张三": 93,
	// 	"李四": 94,
	// }

	scoreMap := make(map[string]int)
	scoreMap["张三"] = 93
	scoreMap["李四"] = 94
	fmt.Println(scoreMap)

	// v表示map中key对应的值，exist表示map中是否存在key
	v, exist := scoreMap["张三"]
	fmt.Println(v, exist)

	for key, value := range scoreMap {
		fmt.Println(key, value)
	}

}
