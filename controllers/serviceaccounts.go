package controllers

import (
	"github.com/gin-gonic/gin"
	"jarvis_server/k8s"
)

func GetServiceaccounts(c *gin.Context) {
	instance, _ := k8s.DefaultManager.K8s("OriginalK8s")
	serviceaccounts := instance.GetSA()

	c.JSON(200, gin.H{
		"code": 20000,
		"data": map[string]interface{}{
			"total": 11111,
			"items":serviceaccounts,
		},
	})
}

func GetNamespacedServiceaccounts(c *gin.Context)  {
	instance, _ := k8s.DefaultManager.K8s("OriginalK8s")
	namespace := c.PostForm("namespace")
	serviceaccounts := instance.GetNamespacedSA(namespace)
	c.JSON(200,gin.H{
		"code":200,
		"data":map[string]interface{}{
			"total": 11111,
			"items":serviceaccounts,
		},
	})
}
