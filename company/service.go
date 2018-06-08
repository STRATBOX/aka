package company

// CompanyService type
type CompanyService struct {
	repo Repository
}

// NewCompanyService returns a new instance of Company service
func NewCompanyService(repo Repository) *CompanyService {
	return &CompanyService{repo}
}

// Find retrives a company for given id
func (s *CompanyService) Find(id UUID) (*Company, error) {
	c, err := s.repo.Find(id)
	if err != nil {
		return nil, err
	}
	return c, err
}

// FindAll returns a list of companies in the database
func (s *CompanyService) FindAll() ([]*Company, error) {
	companies, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}
	return companies, err
}

// Create adds a new company to the database
func (s *CompanyService) Create(c *Company) error {
	err := s.repo.Create(c)
	return err
}

// Update edits a company in the databae with given id
func (s *CompanyService) Update(id UUID, c *Company) error {
	err := s.repo.Update(id, c)
	return err
}

// Delete removes a company in the databae with given id
func (s *CompanyService) Delete(id UUID) error {
	err := s.repo.Delete(id)
	return err
}
