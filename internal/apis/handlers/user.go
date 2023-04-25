package apis

import (
	"log"
	"net/http"

	athena "github.com/dropdevrahul/athena/athena"
	"github.com/dropdevrahul/simple-api/internal/app"
	"github.com/dropdevrahul/simple-api/internal/models"
)

type LoginRequest struct {
	Email    string `json:email`
	Password string `json:password`
}

func UserSignUpHandler(a *app.App, w http.ResponseWriter, r *http.Request) {
	var u models.User

	err := DecodeRequest(w, r, &u)
	if err != nil {
		return
	}

	ut, err := app.SignUpUser(a, &u)
	if err != nil {
		log.Println(err)

		athena.Json(w, models.HTTPError{
			Message: "Failed to save user",
		}, http.StatusInternalServerError, nil)

		return
	}

	athena.Json(w, ut, http.StatusOK, nil)
}

func UserLoginHandler(a *app.App, w http.ResponseWriter, r *http.Request) {
	var req LoginRequest

	err := DecodeRequest(w, r, &req)
	if err != nil {
		return
	}

	ut, err := app.LoginUser(a, req.Email, req.Password)
	if err != nil {
		log.Println(err)

		athena.Json(w, models.HTTPError{
			Message: "Failed to save user",
		}, http.StatusInternalServerError, nil)

		return
	}

	athena.Json(w, ut, http.StatusOK, nil)
}
