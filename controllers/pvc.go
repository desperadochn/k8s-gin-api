package controllers
import (
	"github.com/gin-gonic/gin"
	"jarvis_server/k8s"
)

func PvcList(c *gin.Context) {
	instance, _ := k8s.DefaultManager.K8s("OriginalK8s")
	pvc := instance.GetPvc()

	c.JSON(200, gin.H{
		"code": 20000,
		"data": map[string]interface{}{
			"total": 11111,
			"items": pvc,
		},
	})
}

func NamespacedPvcList(c *gin.Context)  {
	instance, _ := k8s.DefaultManager.K8s("OriginalK8s")
	namespace := c.PostForm("namespace")
	println(namespace)
	pvc := instance.GetNamespacePvc(namespace)
	c.JSON(200, gin.H{
		"code": 20000,
		"data": map[string]interface{}{
			"total": 11111,
			"items": pvc,
		},
	})
}