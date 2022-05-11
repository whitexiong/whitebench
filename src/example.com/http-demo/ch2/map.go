package main


import (
	"fmt"
)

func main() {
	m := make(map[int]string, 9) // 通过 make关键字创建

	// 通过字面量进行创建
	mapA := map[string]int{
		"java": 1,
		"go": 2,
		"php": 3,
	}
	
	m[1] = "hello"
	m[2] = "world"
	m[3] = "go"

	v, ok := m[1]
	_, _ = fn(v, ok)

	delete(m, 1) // 删除 map 中的元素

	// 循环迭代
	for key,value := range mapA{
		fmt.Printf("%s=>%d\n",key,value)
	}

	// for key,value:= range m{
		// fmt.Printf("%s=>%s\n",key,value)
	//  }
}

func fn(v string, ok bool) (string, bool) {
		return v, ok
}
