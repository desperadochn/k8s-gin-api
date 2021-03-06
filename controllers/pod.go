package controllers

import (
	"github.com/gin-gonic/gin"
	"jarvis_server/k8s"
)

func PodList(c *gin.Context) {
	instance, _ := k8s.DefaultManager.K8s("OriginalK8s")
	pod := instance.GetPod()

	c.JSON(200, gin.H{
		"code": 20000,
		"data": map[string]interface{}{
			"total": 11111,
			"items": pod,
		},
	})
}
func NamespacedPodLister(c *gin.Context)  {
	instance, _ := k8s.DefaultManager.K8s("OriginalK8s")
	namespace := c.PostForm("namespace")
	pod := instance.ListNamespacePod(namespace)

	c.JSON(200,gin.H{
		"code":2000,
		"data": map[string]interface{} {
			"total":1111,
			"data": pod,
		},
	})
}