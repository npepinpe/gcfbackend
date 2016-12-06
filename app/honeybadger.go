package app

import (
	"errors"
	"fmt"

	logrus "github.com/Sirupsen/logrus"
	honeybadger "github.com/honeybadger-io/honeybadger-go"
)

type HoneybadgerHook struct {
	App *Application
}

func NewHoneybadgerHook(app *Application) (*HoneybadgerHook, error) {
	hook := HoneybadgerHook{
		App: app,
	}

	return &hook, nil
}

func (hook *HoneybadgerHook) Fire(entry *logrus.Entry) error {
	var err error
	context := make(honeybadger.Context)

	for k, v := range entry.Data {
		if k != logrus.ErrorKey {
			context[k] = v
		}
	}

	err, ok := entry.Data[logrus.ErrorKey].(error)
	if !ok {
		fmt.Println("Could not convert to error")
		err = errors.New(entry.Message)
		honeybadger.Notify(err, honeybadger.ErrorClass{Name: entry.Message}, context)
	} else {
		fmt.Println("Could convert to error")
		honeybadger.Notify(err, context)
	}

	return nil
}

func (hook *HoneybadgerHook) Levels() []logrus.Level {
	return []logrus.Level{logrus.ErrorLevel, logrus.FatalLevel}
}
