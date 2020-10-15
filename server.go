package main

import (
	"SafeToGo/Utils"
	"crypto/tls"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Creates and returns http server of golang's standard library.
// TLS configs are attached to the server.
func createServer(engine *gin.Engine) *http.Server {

	// Configurations for tls connection.
	//
	// TLS_CRT: the SSL/TLS certificate issued by the Certificate Authority
	// TLS_KEY: private key corresponding to the public one in TLS_CRT
	tlsConfig := &tls.Config {

		GetCertificate: certRequestFunc(Utils.GetEnvVar("TLS_CRT"), Utils.GetEnvVar("TLS_KEY")),
	}

	// Http server.
	//
	// Addr: port number that the server should listen to ( set from environment variable )
	// TLSConfig: tls configurations described above
	server := &http.Server {

		Addr:      Utils.GetEnvVar("PORT"),
		Handler:   engine,
		TLSConfig: tlsConfig,
	}

	return server
}

// Designed for tls.Config. Returns function that retrieves tls certificates.
//
// certFile: certificate name to retrieve
// keyFile: private key file name to retrieve
func certRequestFunc(certFile, keyFile string) func(*tls.ClientHelloInfo) (*tls.Certificate, error) {

	return func(hello *tls.ClientHelloInfo) (*tls.Certificate, error) {

		// x.509 is a standard defining the format of public key certificates.
		c, err := tls.LoadX509KeyPair(certFile, keyFile)
		return &c, err
	}
}
