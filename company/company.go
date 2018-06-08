package company

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// UUID type
type UUID string

// Company represents a company/organisation
type Company struct {
	ID          bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	UUID        UUID          `json:"uuid" bson:"uuid"`
	Name        string        `json:"name" bson:"name"`
	DisplayName string        `json:"display.name" bson:"display.name"`
	Description string        `json:"description" bson:"description"`
	Active      bool          `json:"active" bson:"active"`
	Website     string        `json:"website" bson:"website"`
	Ticker      string        `json:"ticker" bson:"ticker"`
	Founded     time.Time     `json:"founded" bson:"founded"`
	Social      social        `json:"social" bson:"social"`
	Sectors     []string      `json:"sectors" bson:"sectors"`
	CreatedAt   time.Time     `json:"createdon" bson:"createdat"`
	UpdatedAt   time.Time     `json:"updatedon" bson:"updatedat"`
}

// social type
type social struct {
	Twitter  string `json:"twitter"`
	Facebook string `json:"facebook"`
	Linkedin string `json:"linkedin"`
}

// Companies is an array of Company
type Companies []*Company

// Repository interface
type Repository interface {
	Find(id UUID) (*Company, error)
	FindAll() ([]*Company, error)
	Create(c *Company) error
	Delete(id UUID) error
	Update(id UUID, c *Company) error
}

// UseCase interface
type UseCase interface {
	Find(id UUID) (*Company, error)
	FindAll() ([]*Company, error)
	Create(c *Company) error
	Delete(id UUID) error
	Update(id UUID, c *Company) error
}
