package main

import (
	"flag"
	"fmt"
	"github.com/sirupsen/logrus"
	"jarvis_server/router"
)

func main() {
	verbose := flag.Bool("v", false, "verbose")
	flag.Usage = func() {
		fmt.Println("usage: server -h")
		flag.PrintDefaults()
	}
	flag.Parse()

	if *verbose {
		// 日志级别为Debug
		logrus.SetLevel(logrus.DebugLevel)
	} else {
		logrus.SetLevel(logrus.InfoLevel)
	}

	router.InitRouter()
}
