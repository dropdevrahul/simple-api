package apis

import (
	"github.com/go-chi/chi"
)

func AddRoutes(
	r *chi.Mux,
) {
	r.Get("/post/{id}", GetPostHandler)
}
