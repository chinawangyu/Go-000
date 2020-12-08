package main

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"log"
	"time"
)

func requestBaidu() string {
	return "baidu content"
}

func requestGoogle() string {
	return "google content"
}

func main() {

	search()

	fmt.Println("处理完成")
	time.Sleep(time.Second * 30)
}

func search() {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	defer cancel()

	g, _ := errgroup.WithContext(ctx)
	g.Go(func() error {
		time.Sleep(time.Second * 2)
		select {
		case <-ctx.Done():
			return errors.New("timeout baidu")
		default:
			response := requestBaidu()
			fmt.Println(response)
		}

		return nil
	})

	g.Go(func() error {
		response := requestGoogle()
		fmt.Println(response)
		return nil
	})

	if err := g.Wait(); err != nil { //等待所有groutine返回, 如果有err返回其中一个
		log.Fatal(err)
	}
	time.Sleep(time.Second*5)
	log.Println("请求完成")
}
