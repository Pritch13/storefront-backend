package api

import (
	"github.com/gin-gonic/gin"
)

func InitializeRoutes(r gin.IRouter) {

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	apiRoutes := r.Group("/api")
	{
		apiRoutes.GET("/items", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "itemshere",
			})
		})
	}

}
