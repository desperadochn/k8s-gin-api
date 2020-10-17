package controllers

import (
	"github.com/gin-gonic/gin"
	"jarvis_server/k8s"
)

func GetRolebindings(c *gin.Context) {
	instance, _ := k8s.DefaultManager.K8s("OriginalK8s")
	rolebindings := instance.GetRolebindings()
	c.JSON(200, gin.H{
		"code": 20000,
		"data": map[string]interface{}{
			"total": 11111,
			"items":rolebindings,
		},
	})
}

func GetNamespacedRolebindings(c *gin.Context)  {
	instance, _ := k8s.DefaultManager.K8s("OriginalK8s")
	namespace := c.PostForm("namespace")
	rolebindings := instance.GetNamespaceRolebindings(namespace)
	c.JSON(200,gin.H{
		"code":200,
		"data":map[string]interface{}{
			"total": 11111,
			"items":rolebindings,
		},
	})
}

