package main

import (
	"fmt"
)


func main(){

	var arr [3]int

	arr[0] = 10
	arr[1] = 11
	arr[2] = 12

	t := arr[0:2]

	fmt.Println(t)
}
