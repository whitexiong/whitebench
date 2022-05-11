package main

import (
	"fmt"
)

func main(){

	var arr [3]int
	// arr[0] = 1
	// fmt.Println(arr[0])
	arr[0] = 999
	arr[1] = 888
	arr[2] = 777

	zero(arr)
	for i, v := range arr{
		fmt.Println("%s ==> %s", i,v)
	}
}



func zero(ptr *[32]byte){ // 32个字节
	for i := range ptr{
		ptr[i] = 0
	}
}