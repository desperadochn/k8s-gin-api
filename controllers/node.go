package controllers

import (
	"github.com/gin-gonic/gin"
	"jarvis_server/k8s"
	_ "jarvis_server/k8s/plugins"
)

func NodeList(c *gin.Context) {
	instance, _ := k8s.DefaultManager.K8s("OriginalK8s")
	node := instance.GetNode()

	c.JSON(200, gin.H{
		"code": 20000,
		"data": map[string]interface{}{
			"total": 11111,
			"items": node,
		},
	})
}
