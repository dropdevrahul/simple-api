package apis

import (
	"encoding/json"
	"log"
	"net/http"

	athena "github.com/dropdevrahul/athena/athena"
	"github.com/dropdevrahul/simple-api/internal/app"
	"github.com/dropdevrahul/simple-api/internal/models"
)


func UserSignUpHandler(a *app.App, w http.ResponseWriter, r *http.Request) {
	var u models.User

	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		log.Println(err)
		athena.Json(w, models.HTTPError{
			Message: "Invalid Payload",
			Errors: []models.ErrorDetail{
				{Detail: err.Error()},
			},
		}, http.StatusBadRequest, nil)

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
