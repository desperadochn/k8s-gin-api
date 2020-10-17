package controllers

import (
	"github.com/gin-gonic/gin"
	"jarvis_server/k8s"
)

func EndpointsLister(c *gin.Context) {
	instance, _ := k8s.DefaultManager.K8s("OriginalK8s")
	endpoints := instance.GetEndpoints()

	c.JSON(200, gin.H{
		"code": 20000,
		"data": map[string]interface{}{
			"total": 11111,
			"items":endpoints,
		},
	})
}

func NamespacedEndpointsLister(c *gin.Context)  {
	instance, _ := k8s.DefaultManager.K8s("OriginalK8s")
	namespace := c.PostForm("namespace")
	endpoints := instance.GetNamespacedEndpoints(namespace)
	c.JSON(200,gin.H{
		"code":200,
		"data":map[string]interface{}{
			"total": 11111,
			"items":endpoints,
		},
	})
}
