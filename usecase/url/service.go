package url

import "github.com/sudo-nick16/shorty/domain/entities"
 
type Service struct {
    repo Repository
}

func NewService(repo Repository) Usecase {
    return &Service{
        repo,
    }
}

func (s *Service) GetURLById(id string) (*entities.URL, error) {
    u, err := s.repo.GetById(id)
    if err != nil {
        return nil, err 
    }
    return u, nil
}

func (s *Service) GetURLByShortURL(shortURL string) (*entities.URL, error) {
    u, err := s.repo.GetByShortURL(shortURL)
    if err != nil {
        return nil, err 
    }
    return u, nil
}

func (s *Service) CreateURL(e *entities.URL) (string, error) {
    u, err := s.repo.Create(e)
    if err != nil {
        return "", err 
    }
    return u, err
}
