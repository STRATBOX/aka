package main

import (
	"log"
	"net/http"
	"time"

	"github.com/STRATBOX/aka/company"
	"github.com/STRATBOX/aka/db"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/koding/multiconfig"

	mgo "gopkg.in/mgo.v2"
)

// const (
// 	// app | service name
// 	appname string = "stratbox.aka.srv.companies"
// )

// Server type for config
type Server struct {
	Port string
}

// Database type for config
type Database struct {
	Name string
	URL  string
}

// Config type for app level settings
type Config struct {
	Server   Server
	Database Database
}

func main() {

	var c Config

	// set config file path directly
	conf := multiconfig.NewWithPath("aka.toml")

	// read config file
	conf.MustLoad(&c)

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
