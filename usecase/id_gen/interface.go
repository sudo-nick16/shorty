package idgen

type  Usecase interface {
    Generate() (string, error) 
}

type ShortIDGenerator interface {
    Generate() (string, error) 
}
