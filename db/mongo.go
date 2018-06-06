package db

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/STRATBOX/aka"
)

// repository type
type repository struct {
	session *mgo.Session
}

// NewRepository creates a new company repository
func NewRepository(session *mgo.Session) *repository {
	return &repository{
		session: session,
	}
}

// Create adds a new company record to mongo
func (r *repository) Create(c *aka.Company) error {
	s := r.session.Clone()
	defer s.Close()

	// insert the company
	err := s.DB("aka").C(collectionam).Insert(c)
	return err
}

// Update amends a company record in mongo
func (r *repository) Update(id string, c *aka.Company) error {
	s := r.session.Clone()
	defer s.Close()

	// update the company
	query := bson.M{"_id": id}
	err := s.DB("aka").C(collectionam).Update(query, c)
	return err
}

// FindByID finds a company record in mongo with id provided
func (r *repository) FindByID(id string) (*aka.Company, error) {
	var c aka.Company
	s := r.session.Clone()
	defer s.Close()

	// find the company
	query := bson.M{"_id": id}
	err := s.DB("aka").C(collectionam).Find(query).One(&c)
	if err != nil {
		return nil, err
	}
	return &c, nil
}

// FindAll returns all companies in the database
func (r *repository) FindAll() (*aka.Companies, error) {
	var companies aka.Companies
	s := r.session.Clone()
	defer s.Close()

	// find the company
	query := bson.M{}
	err := s.DB("aka").C(collectionam).Find(query).All(&companies)
	if err != nil {
		return nil, err
	}
	return &companies, nil
}

// Delete removes a company with given id from the database
func (r *repository) Delete(id string) error {
	s := r.session.Clone()
	defer s.Close()

	// find the company
	query := bson.M{"_id": id}
	err := s.DB("aka").C(collectionam).Remove(query)
	return err
}
