package main

import "fmt"

type req struct {
	reqClueIds []string
}

func main() {

	request := &req{reqClueIds: []string{}}
	fmt.Println(request.reqClueIds == nil)
}
