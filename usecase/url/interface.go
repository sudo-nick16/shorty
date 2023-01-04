package url

import "github.com/sudo-nick16/shorty/domain/entities"

type Repository interface {
	GetById(id string) (*entities.URL, error)
	GetByShortURL(shortURL string) (*entities.URL, error)
	Create(e *entities.URL) (string, error)
}

type Usecase interface {
	GetURLById(id string) (*entities.URL, error)
	GetURLByShortURL(shortURL string) (*entities.URL, error)
	CreateURL(e *entities.URL) (string, error)
}
