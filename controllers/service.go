package controllers

import (
	"github.com/gin-gonic/gin"
	"jarvis_server/k8s"
)

func ServiceList(c *gin.Context) {
	instance, _ := k8s.DefaultManager.K8s("OriginalK8s")
	service := instance.Getservice()

	c.JSON(200, gin.H{
		"code": 20000,
		"data": map[string]interface{}{
			"total": 11111,
			"items":service,
		},
	})
}

func NamespacedService(c *gin.Context)  {
	instance, _ := k8s.DefaultManager.K8s("OriginalK8s")
	namespace := c.PostForm("namespace")
	service := instance.GetNamespacedService(namespace)

	c.JSON(200, gin.H{
		"code": 20000,
		"data": map[string]interface{}{
			"total": 11111,
			"items":service,
		},
	})
}
