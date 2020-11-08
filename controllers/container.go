package controllers

import (
	"github.com/gin-gonic/gin"
	"jarvis_server/k8s"
)

func ContainertLister(c *gin.Context) {
	instance, _ := k8s.DefaultManager.K8s("OriginalK8s")
	container := instance.GetDeploymentContainerName()
	c.JSON(200, gin.H{
		"code": 20000,
		"data": map[string]interface{}{
			"total": 11111,
			"items":container,
		},
	})
}

func NamespaceLabelContainertLister(c *gin.Context)  {
	instance, _ := k8s.DefaultManager.K8s("OriginalK8s")
	//namespace := c.PostForm("namespace")
	//LabelSelector := c.PostForm("labelSelector")
	//labelSelector := LabelSelector
	label := c.PostForm("label")
	namespace := c.PostForm("namespace")
	container := instance.GetLabeldDeploymentContainerName(label,namespace)
	c.JSON(200, gin.H{
		"code": 20000,
		"data": map[string]interface{}{
			"total": 11111,
			"items":container,
		},
	})
}

