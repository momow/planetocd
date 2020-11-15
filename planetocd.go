package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/aureliengasser/planetocd/server"
)

func main() {
	isLocal := isLocalEnvironment()
	var domain string
	var port int

	if isLocal {
		domain = fmt.Sprintf("http://localhost:%v", server.DefaultPort)
		port = server.DefaultPort
	} else {
		domain = "https://" + server.Domain
		port = getPort()
	}

	server.Listen(domain, port)
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
