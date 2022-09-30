package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/golife/gin-layout/config"
	"github.com/golife/gin-layout/database"
	zaplog "github.com/golife/gin-layout/log"
	"github.com/golife/gin-layout/router"
)

func init() {
	path := flag.String("c", "config/settings.yml", "-c config file path")
	flag.Parse()
	config.Init(*path)

	database.Init()
	zaplog.Init()
}

func main() {
	engine := router.Init()
	go func() {
		_ = engine.Run(fmt.Sprintf("%s:%d", config.App().Application.Host, config.App().Application.Port))
	}()
	defer shutdown()

	// 等待中断信号优雅关闭服务器
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
}

func shutdown() {
	if err := database.Close(); err != nil {
		log.Fatal(err)
	}
	log.Println("database closed ...")

	log.Println("server shutdown ...")
}
