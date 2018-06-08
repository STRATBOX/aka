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

	mgo "gopkg.in/mgo.v2"
)

// var session *mgo.Session

const appname string = "aka"
const host string = "localhost"
const port string = ":3333"

// GetMongoSession Creates a new session.
// if mgoSession is nil i.e there is no active mongo session.
// If there is an active mongo session it will return a Clone
func getSession() *mgo.Session {
	session, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		log.Fatal("Failed to start the Mongo session")
	}
	session.SetMode(mgo.Monotonic, true)
	return session
}

func init() {
	db, err := mgo.Dial("mongodb://localhost")
	defer db.Close()

	if err != nil {
		panic(err)
	}

	// fetch collection
	col := db.DB(appname).C("companies")

	index := mgo.Index{
		Key:        []string{"uuid"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}

	err = col.EnsureIndex(index)
	if err != nil {
		panic(err)
	}

	fmt.Printf("[raion]: create database indexes for aka - %s\n", "companies")
}

func main() {
	r := chi.NewRouter()

	// middleware
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	repository := db.NewRepository(getSession())
	companyservice := company.NewService(repository)
	companies := company.NewHandler(companyservice)

	r.Mount("/companies", companies.Routes())

	srv := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		Addr:         ":3000",
		Handler:      r,
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err.Error())
	}

}
