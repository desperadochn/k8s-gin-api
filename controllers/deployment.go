package controllers

import (
	"github.com/gin-gonic/gin"
	"jarvis_server/k8s"
)

func DeploymentLister(c *gin.Context) {
	instance, _ := k8s.DefaultManager.K8s("OriginalK8s")
	deploy := instance.GetDeployment()

	c.JSON(200, gin.H{
		"code": 20000,
		"data": map[string]interface{}{
			"total": 11111,
			"items":deploy,
		},
	})
}

func NamespacedDeploymentLister(c *gin.Context)  {
	instance, _ := k8s.DefaultManager.K8s("OriginalK8s")
	namespace := c.PostForm("namespace")
	deploy := instance.GetNamespacedDeployment(namespace)
	c.JSON(200,gin.H{
		"code":200,
		"data":map[string]interface{}{
			"total": 11111,
			"items":deploy,
		},
	})
}