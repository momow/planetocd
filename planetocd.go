package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/aureliengasser/planetocd/server"
)

func main() {
	isLocal := isLocalEnvironment()
	var scheme string
	var host string
	var port int

	if isLocal {
		scheme = "http"
		host = fmt.Sprintf("localhost:%v", server.DefaultPort)
		port = server.DefaultPort
	} else {
		scheme = "https"
		host = server.Domain
		port = getPort()
	}

	server.Listen(scheme, host, port, isLocal)
}

func isLocalEnvironment() bool {
	environment, ok := os.LookupEnv("ENVIRONMENT")
	if !ok {
		return true
	}

	return environment != "production"
}

func getPort() int {
	portStr, ok := os.LookupEnv("PORT")

	if !ok {
		return server.DefaultPort
	}

	port, err := strconv.Atoi(portStr)

	if err != nil {
		fmt.Printf("Invalid port %s\n", portStr)
		os.Exit(1)
	}

	return port
}
