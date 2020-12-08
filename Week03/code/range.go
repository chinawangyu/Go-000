package main

import "fmt"

func main() {
	//请求 range 取址问题
	slice := []int{1, 2, 3, 4}
	map1 := make(map[int]*int, 4)
	for index, value := range slice {
		map1[index] = &value
	}
	fmt.Printf("%+v", map1)  //4 4 4 4
}
