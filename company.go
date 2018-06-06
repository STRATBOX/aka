package aka

import "time"

// Company represents a company/organisation
type Company struct {
	ID          string    `json:"_id" bson:"_id"`
	Name        string    `json:"name" bson:"name"`
	DisplayName string    `json:"display.name" bson:"display.name"`
	Description string    `json:"description" bson:"description"`
	Active      bool      `json:"active" bson:"active"`
	Website     string    `json:"website" bson:"website"`
	Ticker      string    `json:"ticker" bson:"ticker"`
	Founded     time.Time `json:"founded" bson:"founded"`
	Social      Social    `json:"social" bson:"social"`
	Sectors     []string  `json:"sectors" bson:"sectors"`
	Createdon   time.Time `json:"createdon" bson:"createdon"`
	Updatedon   time.Time `json:"updatedon" bson:"updatedon"`
}

// Social represents social media profiles
type Social struct {
	Twitter  string `json:"twitter" bson:"twitter"`
	Facebook string `json:"facebook" bson:"facebook"`
	Linkedin string `json:"linkedin" bson:"linkedin"`
}

// Companies is an array of Company
type Companies []Company

// Repository interface
type Repository interface {
	Find(id string) (*Company, error)
	FindAll() ([]*Company, error)
	Create(c *Company) error
	Delete(id string) error
	Update(c *Company) error
}

// Service interface
type Service interface {
	Find(id string) (*Company, error)
	FindAll() ([]*Company, error)
	Create(c *Company) error
	Delete(id string) error
	Update(c *Company) error
}
