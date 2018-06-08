package company

// Service type
type Service struct {
	repo Repository
}

// NewService returns a new instance of Company service
func NewService(repo Repository) *Service {
	return &Service{repo}
}

// Find retrives a company for given id
func (s *Service) Find(id UUID) (*Company, error) {
	c, err := s.repo.Find(id)
	if err != nil {
		return nil, err
	}
	return c, err
}

// FindAll returns a list of companies in the database
func (s *Service) FindAll() ([]*Company, error) {
	companies, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}
	return companies, err
}

// Create adds a new company to the database
func (s *Service) Create(c *Company) error {
	err := s.repo.Create(c)
	return err
}

// Update edits a company in the databae with given id
func (s *Service) Update(id UUID, c *Company) error {
	err := s.repo.Update(id, c)
	return err
}

// Delete removes a company in the databae with given id
func (s *Service) Delete(id UUID) error {
	err := s.repo.Delete(id)
	return err
}
