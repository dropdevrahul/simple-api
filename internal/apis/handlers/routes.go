package apis

import (
	"net/http"

	"github.com/dropdevrahul/simple-api/internal/app"
	"github.com/go-chi/chi"
)

func AddRoutes(
	a *app.App,
	r *chi.Mux,
) {
	r.Post("/user", func(w http.ResponseWriter, r *http.Request) {
		UserSignUpHandler(a, w, r)
	})
}
