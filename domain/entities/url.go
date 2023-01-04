package entities

import "errors"

type URL struct {
    Id string `json:"_id" bson:"_id,omitempty"`
    ShortURL string `json:"shortURL,omitempty" bson:"shortURL,omitempty,unique"`
    RedirectTo string `json:"redirectTo,omitempty" bson:"redirectTo,omitempty"`
}

func NewURL(shortURL, redirectTo string) (*URL, error) {
    u := &URL{
        ShortURL: shortURL,
        RedirectTo: redirectTo,
    }
    err := u.Validate()
    if err != nil {
        return nil, err
    }
    return u, nil
}

func (u *URL) Validate() error {
    if u.RedirectTo == "" {
        return errors.New("Few details missing.")
    }
    return nil
}
