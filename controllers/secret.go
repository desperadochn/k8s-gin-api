package controllers

import (
	"github.com/gin-gonic/gin"
	"jarvis_server/k8s"
)

func SecretList(c *gin.Context) {
	instance, _ := k8s.DefaultManager.K8s("OriginalK8s")
	secret := instance.GetSecret()

	c.JSON(200, gin.H{
		"code": 20000,
		"data": map[string]interface{}{
			"total": 11111,
			"items":secret,
		},
	})
}

func NamespacedSecret(c *gin.Context)  {
	instance, _ := k8s.DefaultManager.K8s("OriginalK8s")
	namespace := c.PostForm("namespace")
	secret := instance.GetNamespacedSecret(namespace)

	c.JSON(200, gin.H{
		"code": 20000,
		"data": map[string]interface{}{
			"total": 11111,
			"items":secret,
		},
	})
}
