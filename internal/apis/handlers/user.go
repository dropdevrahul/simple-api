package apis

import (
	"encoding/json"
	"net/http"

	athena "github.com/dropdevrahul/athena/athena"
	"github.com/dropdevrahul/simple-api/internal/application"
	"github.com/dropdevrahul/simple-api/internal/models"
)

type PostResponse struct {
	ID   string `json:"id"`
	Text string `json:"text"`
}

func (a *application.App) UserSignUpHandler(w http.ResponseWriter, r *http.Request) {
	var u models.User
	err := json.NewDecoder(r.Body).Decode(&u)

	athena.Json(w, PostResponse{
		ID:   "12",
		Text: "this is a post",
	}, http.StatusOK, nil)
}
