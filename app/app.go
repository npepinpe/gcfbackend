package app

import (
	"fmt"
	"path"

	"github.com/Sirupsen/logrus"
	"github.com/npepinpe/gcfbackend/utils"
)

type Application struct {
	Root        string
	Config      Configuration
	Environment string
	Logger      *logrus.Logger
	Secrets     map[string]string
}

func NewApplication(rootPath string, environment string) (app *Application, err error) {
	app = &Application{
		Config:      Configuration{},
		Root:        rootPath,
		Environment: environment,
		Logger:      logrus.New(),
	}

	err = app.ReadConfigFile("application.yml", &app.Config, environment)
	if err != nil {
		return
	}

	err = app.ReadConfigFile("secrets.yml", &app.Secrets, environment)
	if err != nil {
		return
	}

	app.Logger.Level, err = logrus.ParseLevel(app.Config.LogLevel)
	if err != nil {
		return
	}
	honeybadgerHook, err := NewHoneybadgerHook(app)
	app.Logger.Hooks.Add(honeybadgerHook)

	return
}

func (app *Application) Path(relativePath string) string {
	return path.Join(app.Root, relativePath)
}

func (app *Application) ReadConfigFile(filename string, config interface{}, environment string) (err error) {
	return utils.ReadEnvironmentYAML(app.Path(fmt.Sprintf("config/%s", filename)), config, environment)
}

type ApplicationError struct {
	Message string
}

func (e ApplicationError) Error() string {
	return e.Message
}
