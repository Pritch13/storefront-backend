package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/pritch13/storefront-backend/api"
)

func main() {
	r := gin.Default()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:3000"}
	corsConfig.AllowCredentials = true

	r.Use(cors.New(corsConfig))
	r.Use(gin.Logger())

	api.InitializeRoutes(r)

	r.Run()
}
