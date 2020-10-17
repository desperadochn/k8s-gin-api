package controllers

import (
	"github.com/gin-gonic/gin"
	"jarvis_server/k8s"
)

func DaemonsetsLister(c *gin.Context) {
	instance, _ := k8s.DefaultManager.K8s("OriginalK8s")
	daemonset := instance.GetDaemonsets()

	c.JSON(200, gin.H{
		"code": 20000,
		"data": map[string]interface{}{
			"total": 11111,
			"items":daemonset,
		},
	})
}

func NamespacedDaemonsetsLister(c *gin.Context)  {
	instance, _ := k8s.DefaultManager.K8s("OriginalK8s")
	namespace := c.PostForm("namespace")
	daemonset := instance.GetNamspacedDaemonsets(namespace)
	c.JSON(200,gin.H{
		"code":200,
		"data":map[string]interface{}{
			"total": 11111,
			"items":daemonset,
		},
	})
}