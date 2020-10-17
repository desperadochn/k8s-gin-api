package controllers

import (
	"github.com/gin-gonic/gin"
	"jarvis_server/k8s"
)

func GetRoles(c *gin.Context) {
	instance, _ := k8s.DefaultManager.K8s("OriginalK8s")
	roles := instance.GetRole()
	c.JSON(200, gin.H{
		"code": 20000,
		"data": map[string]interface{}{
			"total": 11111,
			"items":roles,
		},
	})
}

func GetNamespacedRoles(c *gin.Context)  {
	instance, _ := k8s.DefaultManager.K8s("OriginalK8s")
	namespace := c.PostForm("namespace")
	roles := instance.GetNamespaceRole(namespace)
	c.JSON(200,gin.H{
		"code":200,
		"data":map[string]interface{}{
			"total": 11111,
			"items":roles,
		},
	})
}
