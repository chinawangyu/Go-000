package api

import (
	"fmt"
	"task/service"
)

func GetDataCount() {
	id := 1
	_, err := service.GetServiceDataCount(id)
	if err != nil {
		fmt.Printf("%+v", err) //输出错误
		return
	}

	//todo

	fmt.Println("success")
	return
}
