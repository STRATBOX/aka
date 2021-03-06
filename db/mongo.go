package db

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/STRATBOX/aka/company"
)

// MgoRepository type
type MgoRepository struct {
	db         string
	session    *mgo.Session
	collection string
}

// NewRepository creates a new company repository
func NewRepository(db string, session *mgo.Session) (*MgoRepository, error) {
	r := &MgoRepository{
		db:         db,
		session:    session,
		collection: "companies",
	}

	index := mgo.Index{
		Key:        []string{"uuid"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}

	// fetch collection
	col := session.DB(db).C(r.collection)

	if err := col.EnsureIndex(index); err != nil {
		return nil, err
	}

	return r, nil
}

// Create adds a new company record to mongo
func (r *MgoRepository) Create(c *company.Company) error {
	s := r.session.Clone()
	defer s.Close()

	// insert the company
	// c.ID = bson.NewObjectId()
	err := s.DB(r.db).C(r.collection).Insert(c)
	return err
}

// Find finds a company record in mongo with id provided
func (r *MgoRepository) Find(id company.UUID) (*company.Company, error) {
	var c *company.Company
	s := r.session.Clone()
	defer s.Close()

	// find the company
	// _id := bson.ObjectIdHex(id)
	query := bson.M{"uuid": id}
	err := s.DB(r.db).C(r.collection).Find(query).One(&c)
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
	err := s.DB(r.db).C(r.collection).Find(query).All(&companies)
	if err != nil {
		return nil, err
	}
	return companies, nil
}

// Update amends a company record in mongo
func (r *MgoRepository) Update(id company.UUID, c *company.Company) error {

	s := r.session.Clone()
	defer s.Close()

	// update the company
	query := bson.M{"uuid": id}
	err := s.DB(r.db).C(r.collection).Update(query, c)
	return err
}

// Delete removes a company with given id from the database
func (r *MgoRepository) Delete(id company.UUID) error {
	s := r.session.Clone()
	defer s.Close()

	// find the company
	query := bson.M{"uuid": id}
	err := s.DB(r.db).C(r.collection).Remove(query)
	return err
}
