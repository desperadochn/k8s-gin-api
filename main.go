package main

import (
	"flag"
	"fmt"
	"github.com/sirupsen/logrus"
	"jarvis_server/k8s"
	k8s_ws "jarvis_server/k8s-ws"
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
	var err error

	// 创建k8s客户端
	if k8s_ws.ClientSet, err = k8s_ws.InitClient(); err != nil {
		fmt.Println(err)
		return
	}
	go k8sInit()
	router.InitRouter()
}
