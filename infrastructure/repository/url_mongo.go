package repository

import (
	"context"

	"github.com/sudo-nick16/shorty/domain/entities"
	"github.com/sudo-nick16/shorty/usecase/url"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)


 
type URLMongo struct {
    db *mongo.Database
}

func NewURLMongo(db *mongo.Database) url.Repository {
    return &URLMongo{
        db,
    }
}

func (u *URLMongo) GetById(id string) (*entities.URL, error)  {
    coll := u.db.Collection("urls")
    _id, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return nil, err 
    }
    filter := bson.M{
        "_id": _id, 
    }
    var url entities.URL
    err = coll.FindOne(context.TODO(), filter).Decode(&url)
    if err != nil {
        return nil, err
    }
    return &url, nil
}

func (u *URLMongo) GetByShortURL(shortURL string) (*entities.URL, error)  {
    coll := u.db.Collection("urls")
    filter := bson.M{
        "shortURL": shortURL, 
    }
    var url entities.URL
    err := coll.FindOne(context.TODO(), filter).Decode(&url)
    if err != nil {
        return nil, err
    }
    return &url, nil
}
 
func (u *URLMongo) Create(e *entities.URL) (string, error)  {
    coll := u.db.Collection("urls")
    _, err := coll.InsertOne(context.TODO(), e)
    if err != nil {
        return "", err
    }
    return e.ShortURL, nil
}
