package app

import (
	"database/sql"
	"fmt"
	"path"

	"github.com/Sirupsen/logrus"
	_ "github.com/go-sql-driver/mysql"
	"github.com/npepinpe/gcfbackend/utils"
)

type Application struct {
	Root        string
	Config      Configuration
	Environment string
	Logger      *logrus.Logger
	Secrets     map[string]string
	Db          *sql.DB
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
	if err != nil {
		app.Logger.Warn("Could not configure Honeybadger hook for Logrus")
		err = nil // continue even if the previous step failed
	}

	if len(app.Secrets[ConfigDbPasswordKey]) > 0 {
		app.Config.Database.Password = app.Secrets[ConfigDbPasswordKey]
	}

	if len(app.Secrets[ConfigDbPasswordKey]) > 0 {
		app.Config.Database.User = app.Secrets[ConfigDbUsernameKey]
	}
	app.Db, err = sql.Open("mysql", app.Config.Database.DataSourceName())
	err = app.Db.Ping()
	if err != nil {
		app.Logger.WithError(err).
			Error("Could not open a database connection")
	}

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
