package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pritch13/storefront-backend/api"
)

func main() {
	r := gin.New()

	r.Use(gin.Logger())

	api.InitializeRoutes(r)

	r.Run()
}
