package apis

import (
	"github.com/dropdevrahul/simple-api/internal/application"
	"github.com/go-chi/chi"
)

func AddRoutes(
	a *application.App,
	r *chi.Mux,
) {
	r.Post("/user", a.UserSignUpHandler)
}
