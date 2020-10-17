package controllers

import (
	"github.com/gin-gonic/gin"
	"jarvis_server/k8s"
)

func JobLister(c *gin.Context) {
	instance, _ := k8s.DefaultManager.K8s("OriginalK8s")
	jobs := instance.GetJob()

	c.JSON(200, gin.H{
		"code": 20000,
		"data": map[string]interface{}{
			"total": 11111,
			"items":jobs,
		},
	})
}

func NamespacedJobLister(c *gin.Context)  {
	instance, _ := k8s.DefaultManager.K8s("OriginalK8s")
	namespace := c.PostForm("namespace")
	jobs := instance.GetNamespacedJob(namespace)
	c.JSON(200,gin.H{
		"code":200,
		"data":map[string]interface{}{
			"total": 11111,
			"items":jobs,
		},
	})
}
