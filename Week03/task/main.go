package main


import "fmt"

func main(){

	chanList := make(chan struct{}, 2)
	fmt.Println(len(chanList),cap(chanList)) // 0 2
}