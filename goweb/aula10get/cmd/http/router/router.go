package router

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type router struct {
}

func (router *router) MapRoutes() http.Handler {
	r := chi.NewRouter()

	r.Use(
		middleware.Logger,
		middleware.Timeout(5*time.Second),
	)
	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})
	r.Mount("/", BuildProductRoutes())

	return r
}

func NewRouter() *router {
	return &router{}
}
