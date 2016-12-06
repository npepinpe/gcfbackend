package main

import (
	"encoding/json"
	"net/http"

	"github.com/honeybadger-io/honeybadger-go"
	"github.com/julienschmidt/httprouter"
	"github.com/npepinpe/gcfbackend/app"
)

type Server struct {
	Application *app.Application
	Router      *httprouter.Router
}

func NewServer(application *app.Application) *Server {
	server := Server{
		Application: application,
		Router:      httprouter.New(),
	}
	server.Router.GET("/lbs/beacon_definitions", getDefinitions)

	return &server
}

func (server *Server) Start() {
	address := server.Application.Config.Server.Address()
	server.Application.Logger.Debugf("Starting server at [%s]", address)
	server.Application.Logger.Fatal(http.ListenAndServe(address, honeybadger.Handler(server.Router)))
}

func getDefinitions(response http.ResponseWriter, request *http.Request, params httprouter.Params) {
	data := map[string]interface{}{"a": 1, "b": true}
	renderJSON(response, data)
}

func renderJSON(response http.ResponseWriter, data interface{}) {
	response.Header().Add("Content-Type", "application/json")
	json, err := json.Marshal(data)

	if err != nil {
		panic(err)
	}

	response.Write(json)
}
