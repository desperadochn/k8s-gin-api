package main

import (
	"flag"
	"fmt"
	"github.com/sirupsen/logrus"
	"jarvis_server/k8s"
	"jarvis_server/router"
)

const (
	Host     = "xxxxxx"
	CAData   = "xxxxxx"
	CertData = "xxxxxx"
	KeyData  = "xxxxxx"
)
func k8sInit() {
	instance, ok := k8s.DefaultManager.K8s("OriginalK8s")
	if !ok {
		fmt.Println(ok)
		return
	}
	instance.Init(Host, CAData, CertData, KeyData)
	_ = instance.KubeconfigInit()
}
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
	go k8sInit()
	router.InitRouter()
}
