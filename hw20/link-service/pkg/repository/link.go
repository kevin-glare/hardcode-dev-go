package repository

import (
	"context"
	"github.com/kevin-glare/hardcode-dev-go/hw20/common/pkg/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type LinkRepo struct {
	db *mongo.Collection
}

const CollectionLink = "links"

func NewLinkRepo(db *mongo.Database) *LinkRepo {
	return &LinkRepo{
		db: db.Collection(CollectionLink),
	}
}

func (r *LinkRepo) FindLink(ctx context.Context, query bson.M) (*model.Link, error) {
	var link model.Link
	err := r.db.FindOne(ctx, query).Decode(&link)

	return &link, err
}

func (r *LinkRepo) CreateLink(ctx context.Context, input *model.Link) error {
	_, err := r.db.InsertOne(ctx, input)
	return err
}
