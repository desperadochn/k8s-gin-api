package controllers

import (
	"github.com/gin-gonic/gin"
	"jarvis_server/k8s"
)

func GetClusterrolebinding(c *gin.Context) {
	instance, _ := k8s.DefaultManager.K8s("OriginalK8s")
	clusterrolebinding := instance.GetClusterrolebinding()
	c.JSON(200, gin.H{
		"code": 20000,
		"data": map[string]interface{}{
			"total": 11111,
			"items":clusterrolebinding,
		},
	})
}

