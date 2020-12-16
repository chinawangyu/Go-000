package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"task/http/internal/router"
	"time"
)
var Config config
func Init() {
	/*var configFile = flag.String("f", "./config.toml", "配置文件")
	flag.Parse()
	if _, err := toml.DecodeFile(*configFile, &Config); err != nil {
		panic(err)
	}
*/
	//日志初始化
	//数据库初始化
}

func main() {

	Init()

	g := gin.Default()
	if err := router.Init(g); err != nil {
		log.Fatalf("router.Init %s error", err.Error())
		panic(err)
	}

	srv := &http.Server{
		Addr:           Config.Port,
		Handler:        g,
		ReadTimeout:    20 * time.Second,
		WriteTimeout:   20 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen port %s error", Config.Port)
		}
	}()

	graceShutDown(srv)
}

//销毁退出
func graceShutDown(srv *http.Server) {
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGUSR1, syscall.SIGUSR2)
	<-quit

	log.Println("Shutdown Server ...")

	//创建一个上下文, 10s后超时
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
