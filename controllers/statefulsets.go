package controllers

import (
	"github.com/gin-gonic/gin"
	"jarvis_server/k8s"
)

func StatefulsetsList(c *gin.Context) {
	instance, _ := k8s.DefaultManager.K8s("OriginalK8s")
	statefulsetsList := instance.GetStatefulsets()

	c.JSON(200, gin.H{
		"code": 20000,
		"data": map[string]interface{}{
			"total": 11111,
			"items":statefulsetsList,
		},
	})
}

func NamespacedStatefulsetsList(c *gin.Context)  {
	instance, _ := k8s.DefaultManager.K8s("OriginalK8s")
	namespace := c.PostForm("namespace")
	statefulsetsList := instance.GetNamespacedStatefulsets(namespace)

	c.JSON(200, gin.H{
		"code": 20000,
		"data": map[string]interface{}{
			"total": 11111,
			"items":statefulsetsList,
		},
	})
}