package handler

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (env *Env) hello(
	w http.ResponseWriter,
	r *http.Request,
	_ httprouter.Params) {
	msg, _ := env.dataProvider.Hello()
	sendResposneMessage(w, http.StatusOK, msg)
}
