package aka

// CompanyService type
type service struct {
	repo Repository
}

// NewCompanyService returns a new instance of Company service
func NewCompanyService(repo *Repository) *service {
	return &service{repo}
}

// FindByID retrives a company for given id
func (cs *service) FindByID(id string) (*aka.Company, error) {
	c, err := cs.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	return c, err
}

// FindAll returns a list of companies in the database
func (cs *service) FindAll() (*aka.Companies, error) {
	companies, err := cs.repo.FindAll()
	if err != nil {
		return nil, err
	}
	return companies, err
}

// Create adds a new company to the database
func (cs *service) Create(c *aka.Company) error {
	err := cs.repo.Create(c)
	return err
}

// Update edits a company in the databae with given id
func (cs *service) Update(id string, c *aka.Company) error {
	err := cs.repo.Update(id, c)
	return err
}

// Delete removes a company in the databae with given id
func (cs *service) Delete(id string) error {
	err := cs.repo.Delete(id)
	return err
}
