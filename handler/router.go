package handler

import (
	"go-ws/config"
	"go-ws/data"

	"github.com/julienschmidt/httprouter"
)

type Env struct {
	dataProvider data.DataProvider
	cfg          *config.Config
}

func NewRouter(env *Env) *httprouter.Router {
	router := httprouter.New()

	router.GET("/hello", env.hello)
	router.GET("/secure-hello", env.checkAuth(env.hello))

	return router
}

func NewEnv(cfg *config.Config, dp data.DataProvider) *Env {
	return &Env{
		cfg:          cfg,
		dataProvider: dp,
	}
}
