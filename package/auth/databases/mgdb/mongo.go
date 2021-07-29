package mgdb

import (
	"context"
	"fmt"
	"logit/v1/package/auth"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const DB = "auths"
const UserDB = "users"

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
func (au AuthDb) InsertUser(ctx context.Context, atr *auth.AuthRequest) (interface{}, error) {
	db := ctx.Value("mgClient").(*mongo.Database)

	res, err := db.Collection(UserDB).InsertOne(ctx, bson.D{{Key: "auth_id", Value: atr.ID}})
	// if err != nil && err.Error() == "mongo: no documents in result" {
	// 	res, err := db.Collection(DB).
	// 		InsertOne(ctx, atr)
	// 	return res.InsertedID, err
	// }
	return res.InsertedID, err
}
func (au AuthDb) Update(ctx context.Context, atr *auth.AuthRequest) (interface{}, error) {
	data, err := bson.Marshal(atr)
	if err != nil {
		fmt.Println("48\n", err)
	}
	quey := &bson.M{}
	err = bson.Unmarshal(data, quey)
	if err != nil {
		fmt.Println("53\n", err)
	}
	db := ctx.Value("mgClient").(*mongo.Database)
	res, err := db.Collection(DB).
		UpdateByID(ctx, atr.ID, bson.M{"$set": quey})

	return res.UpsertedID, err
}
