package main

import (
	"SafeToGo/Utils"
	"github.com/gin-gonic/gin"
)

func configureRouter(g *gin.Engine) {

	// Attach paths to the router
	attachMiddlewares(g)
	attachApiPaths(g)
}

func attachMiddlewares(g *gin.Engine) {

	// Write logs to stdout if GIN_MODE is debug
	if Utils.GetEnvVar("GIN_MODE") == "debug" { g.Use(gin.Logger()) }
	// Allowing COR requests from all origins
	g.Use(func(c *gin.Context) {
		referer := c.Request.Referer()
		if len(referer) > 0 && referer[len(referer) - 1] == '/' {
			referer = referer[:len(c.Request.Referer()) - 1]
		}
		c.Writer.Header().Set("Access-Control-Allow-Origin", referer)
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})
	// Auto recover from panics (write internal server error as response)
	g.Use(gin.Recovery())
}

func attachApiPaths(g *gin.Engine) {

	g.GET("/list", getList)
	g.GET("/map", getMap)
}
