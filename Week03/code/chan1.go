package main

import (
	"fmt"
	"sync/atomic"
)

func main() {

	var a int64
	a = 10
	atomic.AddInt64(&a, 12)
	fmt.Println(a)

	//原子性比较交换  如果a=9  则 a+1
	bool := atomic.CompareAndSwapInt64(&a, 9, a+1) //23
	fmt.Println(bool, a)
}
