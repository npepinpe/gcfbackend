package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/honeybadger-io/honeybadger-go"
	"github.com/npepinpe/gcfbackend/app"
)

type Server struct {
	Application *app.Application
	Router      *http.ServeMux
}

func NewServer(application *app.Application) *Server {
	server := Server{
		Application: application,
		Router:      http.NewServeMux(),
	}
	server.Router.HandleFunc("/lbs/beacon_definitions", getDefinitions)

	return &server
}

func (server *Server) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	context := context.WithValue(request.Context(), "application", server.Application)
	server.Router.ServeHTTP(response, request.WithContext(context))
}

func (server *Server) Start() {
	address := server.Application.Config.Server.Address()
	server.Application.Logger.Debugf("Starting server at [%s]", address)
	server.Application.Logger.Fatal(http.ListenAndServe(address, honeybadger.Handler(server)))
}

func getDefinitions(response http.ResponseWriter, request *http.Request) {
	application := request.Context().Value("application").(*app.Application)

	data := map[string]interface{}{"a": 1, "b": true, "environment": application.Environment}
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
