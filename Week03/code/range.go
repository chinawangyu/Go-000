package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//请求 range 取址问题
	slice := []int{1, 2, 3, 4}
	map1 := make(map[int]*int, 4)
	for index, value := range slice {
		map1[index] = &value
	}
	fmt.Printf("%+v", map1)  //4 4 4 4


	//strings count、map
	/*s := "hello wor!ld~@!"
	n := strings.Count(s, "!") //output: 2
	fmt.Println(n)*/

	s := "123test"
	s_map := strings.Map(func(r rune) rune {
		if _, err := strconv.ParseInt(string(r), 0, 64); err != nil {
			return '0'
		}
		return r
	}, s)

	fmt.Println(s_map) //1230000

	//空结构体获取长度
	chanList := make(chan struct{}, 2)
	fmt.Println(len(chanList), cap(chanList)) // 0 2
}
