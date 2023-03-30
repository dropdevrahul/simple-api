BINARY=app

migrate:
	cd migrations && goose postgres "user=${APP_DB_USER} password=${APP_DB_PASSWORD} host=${APP_DB_HOST}\
	 		port=${APP_DB_PORT}	dbname=postgres sslmode=disable" up

migration:
	# name is name of migration passed as argument
	# make migration name=create_some_table
	cd migrations && goose create $(name) sql

build:
	go build -o target/${BINARY}

run:
	go run main.go

# development
dev-db-up:
	cd dev-env && docker compose up -d

dev-db-down:
	cd dev-env && docker compose down
