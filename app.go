package main

import (
	"os"
	"log"
	"fmt"
	"encoding/json"
	"net/http"
	"github.com/emicklei/go-restful"
	restfulspec "github.com/emicklei/go-restful-openapi"
	"github.com/go-openapi/spec"
)

type App struct {
	endpoints []string
	Router *mux.Router
}

func (a *App) Initialize() {
	a.endpoints = os.Args[1:]
	a.Router = mux.NewRouter()
	a.initializeRoutes()
}

func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr , a.Router))
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
    response, _ := json.Marshal(payload)
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(code)
    w.Write(response)
}

func getApiSpec(path string) (json.RawMessage, error) {
	data, err := swag.LoadFromFileOrHTTP(path)
	if err != nil {
		return nil, err
	}
	return json.RawMessage(data), nil
}

func (a *App) aggregateOpenAPISpec(w http.ResponseWriter, r *http.Request) {
	var fullSpec = interface{}
	for _, endpoint := range a.endpoints {
		var spec, _ = getApiSpec(endpoint)

	}
	fmt.Println(fullSpec)
	respondWithJSON(w, http.StatusOK, p)
}

func (a *App) initializeRoutes() {
    a.Router.HandleFunc("/metrics", a.aggregateOpenAPISpec).Methods("GET")
}
