package main

import (
	"github.com/gin-gonic/gin"
)

var engine *gin.Engine

func init() {

	// Creating martini classic server
	engine = gin.New()

	// Configuring gin router's paths and middlewares
	configureRouter(engine)
}

func main() {

	// Creating server from engine (configuring tls)
	server := createServer(engine)

	// Starting the server
	if GetEnvVar("TLS") == "true" {
		// Waiting for tls-encrypted (https) connection
		Must(server.ListenAndServeTLS("", ""))

	} else {
		// Waiting for non-encrypted (http) connection
		Must(server.ListenAndServe())
	}
}
