package controllers

import (
	"github.com/gin-gonic/gin"
	"jarvis_server/k8s"
)

func EvensLister(c *gin.Context) {
	instance, _ := k8s.DefaultManager.K8s("OriginalK8s")
	events := instance.GetEvent()

	c.JSON(200, gin.H{
		"code": 20000,
		"data": map[string]interface{}{
			"total": 11111,
			"items":events,
		},
	})
}

func NamespacedEventsLister(c *gin.Context)  {
	instance, _ := k8s.DefaultManager.K8s("OriginalK8s")
	namespace := c.PostForm("namespace")
	events := instance.GetNamespacedEvent(namespace)
	c.JSON(200,gin.H{
		"code":200,
		"data":map[string]interface{}{
			"total": 11111,
			"items":events,
		},
	})
}
