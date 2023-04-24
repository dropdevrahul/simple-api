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
	"github.com/dropdevrahul/simple-api/internal/app"
	"github.com/dropdevrahul/simple-api/internal/db"
	"github.com/dropdevrahul/simple-api/internal/models"
)

type Server struct {
	App         *app.App
	Injectables map[string]Injectable
	DB          models.DB
	settings    ServerSettings
}

type ServerSettings struct {
	Host string     `mapstrcuture:port`
	Port string     `mapstructure:"port" validate:"required"`
	DB   DBSettings `mapstructure:"db"`
}

type Injectable interface {
}

type DBSettings struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Database string `mapstructure:"database"`
}

func (s *Server) LoadSettings(path string) error {
	config := ServerSettings{}

	viper.AddConfigPath(path)
	viper.SetConfigName("config.local")
	viper.SetConfigType("yml")

	viper.AutomaticEnv()

	envBind := map[string]string{
		"db.host":     "APP_DB_HOST",
		"db.port":     "APP_DB_PORT",
		"db.user":     "APP_DB_USER",
		"db.password": "APP_DB_PASSWORD",
		"db.database": "APP_DB_DB",
	}

	for key, e := range envBind {
		err := viper.BindEnv(key, e)
		if err != nil {
			log.Fatal(err)
			return err
		}
	}

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
		return err
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatal(err)
		return err
	}

	s.settings = config

	return nil
}

func (s *Server) LoadDB() error {
	log.Printf("connecting to db")

	db, err := sqlx.Connect("postgres",
		fmt.Sprintf("host=%s port =%s user=%s password=%s dbname=%s sslmode=disable",
			s.settings.DB.Host, s.settings.DB.Port, s.settings.DB.User, s.settings.DB.Password,
			s.settings.DB.Database,
		))

	if err != nil {
		log.Fatal("Failed connecting to database, ", err)
		return err
	}

	log.Printf("connected to db")
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
	log.Printf("Listening on %s:%s", s.settings.Host, s.settings.Port)
	return http.ListenAndServe(fmt.Sprintf("%s:%s", s.settings.Host, s.settings.Port), r)
}

func (s *Server) LoadApp() {
	s.App = app.NewApp(s.DB, db.NewDBRepo())
}

func RunServer() {
	log.SetFlags(log.LstdFlags | log.Llongfile)

	server := Server{}

	err := server.LoadSettings(".")
	if err != nil {
		panic(err)
	}

	err = server.LoadDB()
	if err != nil {
		panic(err)
	}

	server.LoadApp()

	router := chi.NewRouter()
	//router.Use(middlewares.BasicAuth(func(user, pwd string) error {
	//	return nil
	//}))

	apis.AddRoutes(server.App, router)

	err = server.Serve(router)
	if err != nil {
		panic(err)
	}
}
