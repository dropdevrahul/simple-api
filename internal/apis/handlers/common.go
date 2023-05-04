package apis

import (
	"encoding/json"
	"log"
	"net/http"

	athena "github.com/dropdevrahul/athena/athena"
	"github.com/dropdevrahul/simple-api/internal/models"
)

func DecodeRequest(w http.ResponseWriter,
	r *http.Request, p interface{}) error {
	err := json.NewDecoder(r.Body).Decode(w)
	if err != nil {
		log.Println(err)
		athena.Json(w, models.HTTPError{
			Message: "Invalid Payload",
			Errors: []models.ErrorDetail{
				{Detail: err.Error()},
			},
		}, http.StatusBadRequest, nil)

		return err
	}

	return nil
}
