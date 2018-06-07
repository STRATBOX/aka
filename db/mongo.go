package db

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/STRATBOX/aka/company"
)

// MgoRepository type
type MgoRepository struct {
	session *mgo.Session
}

// NewRepository creates a new company repository
func NewRepository(session *mgo.Session) *MgoRepository {
	return &MgoRepository{
		session: session,
	}
}

// Create adds a new company record to mongo
func (r *MgoRepository) Create(c *company.Company) error {
	s := r.session.Clone()
	defer s.Close()

	// insert the company
	err := s.DB("aka").C("companies").Insert(c)
	return err
}

// Update amends a company record in mongo
func (r *MgoRepository) Update(id string, c *company.Company) error {

	s := r.session.Clone()
	defer s.Close()

	// update the company
	_id := bson.ObjectIdHex(id)
	query := bson.M{"_id": _id}
	err := s.DB("aka").C("companies").Update(query, c)
	return err
}

// Find finds a company record in mongo with id provided
func (r *MgoRepository) Find(id string) (*company.Company, error) {
	var c *company.Company
	s := r.session.Clone()
	defer s.Close()

	// find the company
	_id := bson.ObjectIdHex(id)
	query := bson.M{"_id": _id}
	err := s.DB("aka").C("companies").Find(query).One(&c)
	if err != nil {
		return nil, err
	}
	return c, nil
}

// FindAll returns all companies in the database
func (r *MgoRepository) FindAll() ([]*company.Company, error) {
	var companies []*company.Company
	s := r.session.Clone()
	defer s.Close()

	// find the company
	query := bson.M{}
	err := s.DB("aka").C("companies").Find(query).All(&companies)
	if err != nil {
		return nil, err
	}
	return companies, nil
}

// Delete removes a company with given id from the database
func (r *MgoRepository) Delete(id string) error {
	s := r.session.Clone()
	defer s.Close()

	// find the company
	_id := bson.ObjectIdHex(id)
	query := bson.M{"_id": _id}
	err := s.DB("aka").C("companies").Remove(query)
	return err
}
