package apiserver

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"

	apis "github.com/dropdevrahul/simple-api/internal/apis/handlers"
	"github.com/dropdevrahul/simple-api/internal/application"
	"github.com/dropdevrahul/simple-api/internal/db"
	"github.com/dropdevrahul/simple-api/internal/middlewares"
	"github.com/dropdevrahul/simple-api/internal/models"
)

type Server struct {
	App         application.App
	Injectables map[string]Injectable
	DB          models.DB
	settings    ServerSettings
}

type ServerSettings struct {
	Port string     `mapstructure:"port" validate:"required"`
	db   DBSettings `mapstructure:"db"`
}

type Injectable interface {
}

type DBSettings struct {
	host     string `mapstructure:"host"`
	port     string `mapstructure:"port"`
	user     string `mapstructure:"user"`
	password string `mapstructure:"password"`
	database string `mapstructure:"database"`
}

func (s *Server) LoadSettings(path string) error {
	config := ServerSettings{}

	viper.AddConfigPath(path)
	viper.SetConfigName("config.local")
	viper.SetConfigType("yml")

	viper.AutomaticEnv()

	err := viper.BindEnv("db.host", "APP_DB_HOST")
	err = viper.BindEnv("db.port", "APP_DB_PORT")
	err = viper.BindEnv("db.database", "APP_DB")
	err = viper.BindEnv("db.user", "APP_DB_USER")
	err = viper.BindEnv("db.password", "APP_DB_PASSWORD")
	if err != nil {
		return err
	}

	err = viper.ReadInConfig()
	if err != nil {
		return err
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return err
	}

	log.Printf("%+v", config)

	s.settings = config

	return nil
}

func (s *Server) LoadDB() error {
	log.Printf("connecting to db %s:%s", s.settings.db.host, s.settings.db.port)

	db, err := sqlx.Connect("postgres",
		fmt.Sprintf("host=%s port =%s user=%s password=%s dbname=%s sslmode=disable",
			s.settings.db.host, s.settings.db.port, s.settings.db.user, s.settings.db.password,
			s.settings.db.database,
		))

	if err != nil {
		log.Fatal(err)
	}

	s.DB = models.DB{
		DB: db,
	}

	return nil
}

func (s *Server) Inject(key string, injectable Injectable) error {
	s.Injectables[key] = injectable
	return nil
}

func (s *Server) Serve(r *chi.Mux) error {
	return http.ListenAndServe(s.settings.Port, r)
}

func (s *Server) LoadApp() error {
	s.App := application.NewApp(s.DB, db.NewDBRepo())
}

func RunServer() {
	server := Server{
		Settings: ServerSettings{
			Port: ":8080",
		},
	}

	err := server.LoadSettings(".")
	if err != nil {
		panic(err)
	}

	err = server.LoadDB()
	if err != nil {
		panic(err)
	}

	s.LoadApp()

	router := chi.NewRouter()
	router.Use(middlewares.BasicAuth(func(user, pwd string) error {
		return nil
	}))

	apis.AddRoutes(&s.App, router)

	err = server.Serve(router)
	if err != nil {
		panic(err)
	}
}
