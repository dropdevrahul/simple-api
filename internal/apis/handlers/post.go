package apis

import (
	"net/http"

	athena "github.com/dropdevrahul/athena/athena"
)

type PostResponse struct {
	ID   string `json:"id"`
	Text string `json:"text"`
}

func GetPostHandler(w http.ResponseWriter, r *http.Request) {
	athena.Json(w, PostResponse{
		ID:   "12",
		Text: "this is a post",
	}, http.StatusOK, nil)
}
