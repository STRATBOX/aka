package aka

import "time"

// Company represents a company/organisation
type Company struct {
	Name        string    `json:"name" bson:"name"`
	DisplayName string    `json:"display.name" bson:"display.name"`
	Description string    `json:"description" bson:"description"`
	Active      bool      `json:"active" bson:"active"`
	Website     string    `json:"website" bson:"website"`
	Ticker      string    `json:"ticker" bson:"ticker"`
	Founded     time.Time `json:"founded" bson:"founded"`
	Social      Social    `json:"social" bson:"social"`
	Sectors     []string  `json:"sectors" bson:"sectors"`
	CreatedAt   time.Time `json:"createdon" bson:"createdat"`
	UpdatedAt   time.Time `json:"updatedon" bson:"updatedat"`
}

// Social represents social media profiles
type Social struct {
	Twitter  string `json:"twitter" bson:"twitter"`
	Facebook string `json:"facebook" bson:"facebook"`
	Linkedin string `json:"linkedin" bson:"linkedin"`
}

// Companies is an array of Company
type Companies []*Company

// Repository interface
type Repository interface {
	Find(id string) (*Company, error)
	FindAll() ([]*Company, error)
	Create(c *Company) error
	Delete(id string) error
	Update(id string, c *Company) error
}

// Service interface
type Service interface {
	Find(id string) (*Company, error)
	FindAll() ([]*Company, error)
	Create(c *Company) error
	Delete(id string) error
	Update(id string, c *Company) error
}
