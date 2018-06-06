package main

import (
	"fmt"
	"log"

	mgo "gopkg.in/mgo.v2"
)

var session *mgo.Session

// GetMongoSession Creates a new session.
// if mgoSession is nil i.e there is no active mongo session.
// If there is an active mongo session it will return a Clone
func GetMongoSession() *mgo.Session {
	if session == nil {
		var err error
		session, err = mgo.Dial("mongodb://localhost")
		if err != nil {
			log.Fatal("Failed to start the Mongo session")
		}
	}
	session.SetMode(mgo.Monotonic, true)
	return session.Clone()
}

func main() {
	fmt.Printf("Running....")
}
