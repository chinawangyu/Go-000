package main

import (
	"context"
	"golang.org/x/sync/errgroup"
	"log"
)

func main() {

	g, _ := errgroup.WithContext(context.Background())
	g.Go(func() error {
		registerAppServer()
		return nil
	})

	g.Go(func() error {
		registerSignal()
		return nil
	})
	
	if err := g.Wait(); err != nil {
		log.Fatal("启动服务失败", err.Error())
	}
}

func registerAppServer() {

}

func registerSignal() {

}
