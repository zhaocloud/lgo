package lgo

import (
	"fmt"
	"net/http"
	"runtime/debug"
)

type Middleware interface {
	WrapFunc(handler HandlerFunc) HandlerFunc
}

func WrapMiddlewares(middlewares []Middleware, handler HandlerFunc) HandlerFunc {
	wrapped := handler
	for i := len(middlewares) - 1; i >= 0; i-- {
		wrapped = middlewares[i].WrapFunc(wrapped)
	}
	return wrapped
}

type errorMiddleware struct{}

func (m errorMiddleware) WrapFunc(handler HandlerFunc) HandlerFunc {
	return func(w ResponseWriter, r *Request) {
		// catch user code's panic, and convert to http response
		defer func() {
			if reco := recover(); reco != nil {
				trace := debug.Stack()
				message := fmt.Sprintf("%s\n%s", reco, trace)
				Abort(w, message, http.StatusInternalServerError)
			}
		}()
		handler(w, r)
	}
}
