package main

import (
	"fmt"
	"go-ws/config"
	"go-ws/data"
	"go-ws/handler"
	"log"
	"net/http"
)

func main() {
	cfg, err := config.GetConfigFromFile("config.json")
	if err != nil {
		log.Fatal("Config parsing error: ", err)
	}

	dataProvider, err := data.NewDataProvider(cfg)
	if err != nil {
		log.Fatal("Init service provider error: ", err)
	}

	env := handler.NewEnv(cfg, dataProvider)

	router := handler.NewRouter(env)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", cfg.Port), router))
}
