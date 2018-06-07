package company

import (
	"time"
)

// Company represents a company/organisation
type Company struct {
	ID          ID        `json:"id" bson:"_id"`
	UUID        string    `json:"uuid" bson:"uuid"`
	Name        string    `json:"name" bson:"name"`
	DisplayName string    `json:"display.name" bson:"display.name"`
	Description string    `json:"description" bson:"description"`
	Active      bool      `json:"active" bson:"active"`
	Website     string    `json:"website" bson:"website"`
	Ticker      string    `json:"ticker" bson:"ticker"`
	Founded     time.Time `json:"founded" bson:"founded"`
	Social      social    `json:"social" bson:"social"`
	Sectors     []string  `json:"sectors" bson:"sectors"`
	CreatedAt   time.Time `json:"createdon" bson:"createdat"`
	UpdatedAt   time.Time `json:"updatedon" bson:"updatedat"`
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
	Find(id ID) (*Company, error)
	FindAll() ([]*Company, error)
	Create(c *Company) error
	Delete(id ID) error
	Update(id ID, c *Company) error
}

// Service interface
type Service interface {
	Find(id ID) (*Company, error)
	FindAll() ([]*Company, error)
	Create(c *Company) error
	Delete(id ID) error
	Update(id ID, c *Company) error
}
