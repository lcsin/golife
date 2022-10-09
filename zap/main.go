package main

import "github.com/golife/zap/pkg"

func init() {
	pkg.Init()
}

func main() {
	pkg.Info("test zap pkg log ...")
}
