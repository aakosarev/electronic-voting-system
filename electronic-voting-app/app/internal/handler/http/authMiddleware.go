package http

import (
	"context"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func AuthMiddleware(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		c, err := r.Cookie("session_token")
		if err != nil {
			if err == http.ErrNoCookie {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		sessionToken := c.Value

		userSession, exists := sessions[sessionToken]
		if !exists {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		if userSession.isExpired() {
			delete(sessions, sessionToken)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "cookie_session_token", c)

		next(w, r.WithContext(ctx), ps)
	}
}
