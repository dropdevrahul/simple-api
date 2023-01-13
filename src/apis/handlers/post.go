package apis

import (
	"net/http"

	zeus "github.com/dropdevrahul/zeus/src"
)

type PostResponse struct {
	ID   string `json:"id"`
	Text string `json:"text"`
}

func GetPostHandler(w http.ResponseWriter, r *http.Request) {
	zeus.Json(w, PostResponse{
		ID:   "12",
		Text: "this is a post",
	}, http.StatusOK, nil)
}
