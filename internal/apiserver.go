package apiserver

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

type Server struct {
	Injectables map[string]Injectable
	Db          Db
	Settings    ServerSettings
}

type ServerSettings struct {
	Port string     `mapstructure:"port" validate:"required"`
	DB   DBSettings `mapstructure:"db"`
}

type DBSettings struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Database string `mapstructure:"database"`
}

type Injectable interface {
}

type Db struct {
	DB *sqlx.DB
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

	s.Settings = config

	return nil
}

func (s *Server) LoadDB() error {
	log.Printf("connecting to db %s:%s", s.Settings.DB.Host, s.Settings.DB.Port)

	db, err := sqlx.Connect("postgres",
		fmt.Sprintf("host=%s port =%s user=%s password=%s dbname=%s sslmode=disable",
			s.Settings.DB.Host, s.Settings.DB.Port, s.Settings.DB.User, s.Settings.DB.Password,
			s.Settings.DB.Database,
		))

	if err != nil {
		log.Fatal(err)
	}

	s.Db = Db{
		DB: db,
	}

	return nil
}

func (s *Server) Inject(key string, injectable Injectable) error {
	s.Injectables[key] = injectable
	return nil
}

func (s *Server) Serve(r *chi.Mux) error {
	return http.ListenAndServe(s.Settings.Port, r)
}
