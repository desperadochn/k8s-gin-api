package controllers

import (
	"github.com/gin-gonic/gin"
	"jarvis_server/k8s"
)

func ConfigMapList(c *gin.Context) {
	instance, _ := k8s.DefaultManager.K8s("OriginalK8s")
	cm := instance.GetConfigMap()

	c.JSON(200, gin.H{
		"code": 20000,
		"data": map[string]interface{}{
			"total": 11111,
			"items":cm,
		},
	})
}

func NamespacedConfigmap(c *gin.Context)  {
	instance, _ := k8s.DefaultManager.K8s("OriginalK8s")
	namespace := c.PostForm("namespace")
	cm := instance.GetNamespacedConfigMap(namespace)

	c.JSON(200, gin.H{
		"code": 20000,
		"data": map[string]interface{}{
			"total": 11111,
			"items":cm,
		},
	})
}
