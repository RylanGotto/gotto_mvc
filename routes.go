package tonic

import (
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func (t *Tonic) routes() http.Handler {
	mux := chi.NewRouter()
	mux.Use(middleware.RequestID)
	mux.Use(middleware.RealIP)
	if t.Debug {
		mux.Use(middleware.Logger)
	}
	mux.Use(middleware.Recoverer)

	return mux
}
