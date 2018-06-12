package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/STRATBOX/aka/company"
	"github.com/STRATBOX/aka/db"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/spf13/viper"

	mgo "gopkg.in/mgo.v2"
)

// const (
// 	// app | service name
// 	appname string = "stratbox.aka.srv.companies"
// )

func main() {
	// create config struct
	type config struct {
		Server struct {
			Port string `mapstructure:"port"`
		} `mapstructure:"server"`
		Database struct {
			Name string `mapstructure:"name"`
			URL  string `mapstructure:"url"`
		} `mapstructure:"database"`
	}

	// load environment variables
	var c config

	// set config file path directly
	// viper.SetConfigFile("aka.json")
	// Add paths config paths. Accepts multiple paths.
	// It will search these paths in given order
	viper.AddConfigPath(".")
	// viper.AddConfigPath("./config")
	// register config filename (no extension)
	viper.SetConfigName("aka")
	// optionally set confilg type
	viper.SetConfigType("toml")

	// read config file
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	// unmarshal config into struct
	if err := viper.Unmarshal(&c); err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}
	fmt.Printf("port=%s db=%s url=%s\n", c.Server.Port, c.Database.Name, c.Database.URL)

	r := chi.NewRouter()

	// middleware
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	repository, _ := db.NewRepository(c.Database.Name, getSession(c.Database.URL))
	companyservice := company.NewService(repository)
	companies := company.NewHandler(companyservice)

	r.Mount("/companies", companies.Routes())

	srv := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		Addr:         c.Server.Port,
		Handler:      r,
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err.Error())
	}

}

// GetMongoSession Creates a new session.
// if mgoSession is nil i.e there is no active mongo session.
// If there is an active mongo session it will return a Clone
func getSession(url string) *mgo.Session {
	session, err := mgo.Dial(url)
	if err != nil {
		log.Fatal("Failed to start the Mongo session")
	}
	session.SetMode(mgo.Monotonic, true)
	return session
}
