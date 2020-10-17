
package controllers

import (
"github.com/gin-gonic/gin"
"jarvis_server/k8s"
)

func StorageclassesList(c *gin.Context) {
	instance, _ := k8s.DefaultManager.K8s("OriginalK8s")
	storageclasses := instance.GetStorageclasses()

	c.JSON(200, gin.H{
		"code": 20000,
		"data": map[string]interface{}{
			"total": 11111,
			"items":storageclasses,
		},
	})
}


