package controllers

import "github.com/gin-gonic/gin"

func Login(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": 20000,
		"data": "admin-token",
	})
}

func Info(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": 20000,
		"data": map[string]interface{}{
			"roles":        []string{"admin"},
			"introduction": "I am a super administrator",
			"avatar":       "https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif",
			"name":         "Super Admin",
		},
	})
}
