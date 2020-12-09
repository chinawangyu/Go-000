package main

import (
	"context"
	"errors"
	"golang.org/x/sync/errgroup"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var done = make(chan struct{}, 1)
var srv = &http.Server{Addr: ":8888"}

//注册http服务
func registerHttpServer() error {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		io.WriteString(writer, "hello world")
	})
	if err := srv.ListenAndServe(); err != nil {
		done <- struct{}{}
		return err
	}
	return nil
}

//注册信号
func registerSignal() error {
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-quit:
		done <- struct{}{}
	case <-done:
		return errors.New("结束服务")
	}
	return nil
}

//关闭服务
func shutDownServer() error {
	<-done
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
		return err
	}
	return nil
}

func main() {

	g, _ := errgroup.WithContext(context.Background())
	//启动 http 服务
	g.Go(func() error {
		if err := registerHttpServer(); err != nil {
			return err
		}
		return nil
	})

	//注册信号
	g.Go(func() error {
		if err := registerSignal(); err != nil {
			return err
		}
		return nil
	})

	//清理工作
	g.Go(func() error {
		if err := shutDownServer(); err != nil {
			return err
		}
		return nil
	})

	if err := g.Wait(); err != nil {
		log.Fatal(err.Error())
	}
}
