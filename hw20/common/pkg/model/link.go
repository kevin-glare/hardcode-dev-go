package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Link struct {
	ID       primitive.ObjectID `json:"-"        bson:"_id,omitempty"`
	Url      string             `json:"url"       bson:"url"`
	ShortUrl string             `json:"short_url" bson:"short_url"`
}
