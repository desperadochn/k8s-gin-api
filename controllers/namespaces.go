package controllers

import (
	"github.com/gin-gonic/gin"
	"jarvis_server/models"
)

func NameSpacesList(c *gin.Context) {

	c.JSON(200, gin.H{
		"code": 20000,
		"data": map[string]interface{}{
			"total": 11111,
			"items": models.QueryNameSpacesAll(),
		},
	})
}
