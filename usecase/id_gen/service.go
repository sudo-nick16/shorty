package idgen

type Service struct {
    idGen ShortIDGenerator
}

func NewService(idGen ShortIDGenerator) Usecase {
    return &Service{
        idGen,
    }
}

func (s *Service) Generate() (string, error) {
    shortId, err := s.idGen.Generate()
    if err != nil {
        return "", err 
    }
    return shortId, nil 
}
