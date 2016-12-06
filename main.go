package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/honeybadger-io/honeybadger-go"
	"github.com/npepinpe/gcfbackend/app"
)

func init() {
}

func main() {
	application, err := initializeApplication()

	if err != nil {
		panic(err)
	}

	application.Logger.Debugf("Configured application with environment [%s] and root dir [%s]", application.Environment, application.Root)
	application.Logger.Debugf("Connected to DB [%s]", application.Config.Database.Dbname)

	// Setup Honeybadger
	configureHoneybadger(application)
	defer honeybadger.Monitor()
	application.Logger.Debugf("Configured Honeybadger monitoring to send errors: %t", application.Config.Honeybadger.Enabled)

	// Start server
	server := NewServer(application)
	server.Start()
}

func initializeApplication() (application *app.Application, err error) {
	var rootPath string
	var environment string
	cwd, err := os.Getwd()

	if err != nil {
		return
	}

	flag.StringVar(&rootPath, "r", cwd, fmt.Sprintf("path to the configuration directory (default: %s)", cwd))
	flag.StringVar(&environment, "e", "development", "current environment (default: development)")
	flag.Parse()

	rootPath = strings.TrimSpace(rootPath)
	if len(rootPath) == 0 {
		rootPath = cwd
	}

	environment = strings.TrimSpace(environment)
	if len(environment) == 0 {
		environment = "development"
	}

	application, err = app.NewApplication(rootPath, environment)
	return
}

func configureHoneybadger(application *app.Application) {
	configuration := honeybadger.Configuration{
		APIKey: application.Secrets[app.ConfigHoneybadgerAPIKey],
		Env:    application.Environment,
		Root:   application.Root,
	}

	if !application.Config.Honeybadger.Enabled {
		configuration.Backend = honeybadger.NewNullBackend()
	}

	honeybadger.Configure(configuration)
}
