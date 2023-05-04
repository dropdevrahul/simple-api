# Go API
A simple go API with minimum dependencies and postgres integration


## Dependencies 

- Database
  - [sqlx](https://github.com/jmoiron/sqlx) to connect to db 
  - [goose](https://github.com/pressly/goose) for migration
  - [lib-pq](https://github.com/lib/pq) as database driver

- Configuration
  - [viper](https://github.com/spf13/viper) for configuration

- Router
  - [chi](https://github.com/go-chi/chi) as router

- API Docs
  - [swag](https://github.com/swaggo/swag#getting-started) helps generating swagger from docs as well as serve swagger UI from within the API using any popular routers

## Makefile

Please see the makefile for available commands.


- To generate swagger API use the make command `make gen-swagger` then run `make run` and go to [localhost](http://localhost:9090/swagger/index.html) to see 
the generated swagger docs 
