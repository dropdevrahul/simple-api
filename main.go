package main

import (
	apiserver "github.com/dropdevrahul/simple-http-server/apiserver/src"
	apis "github.com/dropdevrahul/simple-http-server/apiserver/src/apis/handlers"
	"github.com/dropdevrahul/simple-http-server/apiserver/src/middlewares"
	"github.com/go-chi/chi"
)

func main() {
	server := apiserver.Server{
		Settings: apiserver.ServerSettings{
			Port: ":8080",
		},
	}
	router := chi.NewRouter()
	router.Use(middlewares.BasicAuth(func(user, pwd string) error {
		return nil
	}))

	apis.AddRoutes(router)

	_ = server.LoadDB()

	server.Serve(router)
}
