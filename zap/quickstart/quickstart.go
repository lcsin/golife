package main

import (
	"fmt"
	"time"

	"go.uber.org/zap"
)

func useLogger() {
	log1 := zap.NewExample()
	log2, _ := zap.NewDevelopment()
	log3, _ := zap.NewProduction()

	log1.Info("log1 quickstart ...")
	log2.Info("log2 quickstart ...")
	log3.Info("log3 quickstart ...")
}

func useSugarLogger() {
	log1 := zap.NewExample()
	log2, _ := zap.NewDevelopment()
	log3, _ := zap.NewProduction()

	sugar1 := log1.Sugar()
	sugar2 := log2.Sugar()
	sugar3 := log3.Sugar()

	sugar1.Infof("sugar1: %v", "quickstart ...")
	sugar2.Infof("sugar2: %v", "quickstart ...")
	sugar3.Infof("sugar3: %v", "quickstart ...")
}

func main() {
	useLogger()

	time.Sleep(1 * time.Second)
	fmt.Println("=========")

	useSugarLogger()
}
