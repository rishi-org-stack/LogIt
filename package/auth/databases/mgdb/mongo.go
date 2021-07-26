package mgdb

import (
	"context"
	"logit/v1/package/auth"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const DB = "auths"

type AuthDb struct{}

func (auth AuthDb) ConnectToCollection() {

}

func (au AuthDb) FindOrInsert(ctx context.Context, atr *auth.AuthRequest) (interface{}, error) {
	var authReq = &auth.AuthRequest{}
	db := ctx.Value("mgClient").(*mongo.Database)
	err := db.Collection(DB).
		FindOne(ctx, bson.D{{
			"email", atr.Email,
		}}).
		Decode(authReq)
	if err != nil && err.Error() == "mongo: no documents in result" {
		res, err := db.Collection(DB).
			InsertOne(ctx, atr)
		return res.InsertedID, err
	}
	return authReq, nil
}
