package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"
)

type resultStruct struct {
	Data string
	Code int
}

//超时退出
func main() {
	if err := search2(); err != nil {
		log.Fatal(err.Error())
	}
}

func search2() error {

	//定义超时时间 5 s
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	ch := make(chan resultStruct)

	go func() {
		//运行6s
		time.Sleep(time.Second * 6)
		result := resultStruct{
			Data: "hello world",
			Code: 2000,
		}
		ch <- result
	}()

	select {
	case <-ctx.Done():
		return errors.New("超时退出")
	case result := <-ch:
		if result.Code != 2000 {
			return errors.New("出现错误")
		}
		fmt.Println(result.Data)
		return nil
	}
}
