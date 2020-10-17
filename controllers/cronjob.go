package controllers

import (
	"github.com/gin-gonic/gin"
	"jarvis_server/k8s"
)

func CronJobLister(c *gin.Context) {
	instance, _ := k8s.DefaultManager.K8s("OriginalK8s")
	Cronjobs := instance.GetCronJob()

	c.JSON(200, gin.H{
		"code": 20000,
		"data": map[string]interface{}{
			"total": 11111,
			"items":Cronjobs,
		},
	})
}

func NamespacedCronJobLister(c *gin.Context)  {
	instance, _ := k8s.DefaultManager.K8s("OriginalK8s")
	namespace := c.PostForm("namespace")
	Cronjobs := instance.GetNamespacedCronjob(namespace)
	c.JSON(200,gin.H{
		"code":200,
		"data":map[string]interface{}{
			"total": 11111,
			"items":Cronjobs,
		},
	})
}
