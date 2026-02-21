package notes

type Service struct { repo *Repository }

func NewService(repo *Repository) *Service { return &Service{repo: repo} }

func (s *Service) CreateNote(title, content string) (*Note, error) {
	n := &Note{Title: title, Content: content}
	return n, s.repo.Create(n)
}

func (s *Service) GetNotes() ([]Note, error) { return s.repo.GetAll() }
func (s *Service) GetNote(id int) (*Note, error) { return s.repo.GetByID(id) }
func (s *Service) UpdateNote(id int, title, content string) (*Note, error) {
	n, err := s.repo.GetByID(id)
	if err != nil { return nil, err }
	n.Title, n.Content = title, content
	return n, s.repo.Update(n)
}
func (s *Service) DeleteNote(id int) error { return s.repo.Delete(id) }