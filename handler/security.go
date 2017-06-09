package handler

import (
	"context"
	"go-ws/util"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

const CTX_USER_ID = 0

func (env *Env) checkAuth(h httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		token := r.Header.Get("access-token")
		if token == "" {
			sendResposneMessage(w, http.StatusUnauthorized, "Missing token")
			return
		}
		claims, err := util.ParseToken(token, env.cfg.TokenSecret)
		if err != nil {
			if err == util.ErrorTokenExpired {
				sendResposneMessage(w, http.StatusUnauthorized, "Token expired")
			} else {
				sendResposneMessage(w, http.StatusUnauthorized, "Token invalid")
			}
			return
		}

		ctx := context.WithValue(r.Context(), CTX_USER_ID, claims.Id)
		r = r.WithContext(ctx)
		h(w, r, p)
	}
}

func ctxGetUsername(ctx context.Context) int64 {
	return ctx.Value(CTX_USER_ID).(int64)
}
