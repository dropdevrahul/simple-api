package main

import (
	apiserver "github.com/dropdevrahul/simple-http-server/apiserver/internal"
	apis "github.com/dropdevrahul/simple-http-server/apiserver/internal/apis/handlers"
	"github.com/dropdevrahul/simple-http-server/apiserver/internal/middlewares"
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

	err := server.LoadSettings(".")
	if err != nil {
		panic(err)
	}

	err = server.LoadDB()

	if err != nil {
		panic(err)
	}

	server.Serve(router)
}
