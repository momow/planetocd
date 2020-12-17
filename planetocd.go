package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/urfave/cli/v2"

	"github.com/aureliengasser/planetocd/server"
	"github.com/aureliengasser/planetocd/translate"
)

func main() {
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "action",
				Required: true,
			},
		},
		Action: func(c *cli.Context) error {
			action := c.String("action")
			switch action {
			case "server":
				startServer()
			case "convert":
				convert()
			case "translate":
				// GOOGLE_APPLICATION_CREDENTIALS=~/.local/planetocd-86fb09efe9c9.json
				doTranslate()
			default:
				log.Fatalf("Unknown action %v", action)
			}
			return nil
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func startServer() {
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

func convert() {
	translate.HTMLToMarkdown()
}

func doTranslate() {
	translate.CreateTranslatedArticle(
		"0003",
		"https://www.madeofmillions.com/articles/12-signs-that-you-might-have-sexual-orientation-ocd",
		"Dr. Jan Weiner",
		"12 Signs That You Might Have Sexual Orientation OCD")
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
