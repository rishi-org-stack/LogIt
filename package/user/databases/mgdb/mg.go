package mgdb

import (
	"context"
	"fmt"
	user "logit/v1/package/user"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const UserDB = "users"

type UserDb struct{}

func (Utb *UserDb) GetUser(ctx context.Context, id string) ([]user.UserAggregate, error) {
	var res = []user.UserAggregate{}
	// Id, _ := primitive.ObjectIDFromHex(id)
	db := ctx.Value("mgClient").(*mongo.Database)
	// err := db.Collection(UserDB).
	// 	FindOne(ctx, bson.D{{"_id", Id}}).Decode(res)
	lookup := bson.D{{
		Key: "$lookup",
		Value: bson.D{{
			Key:   "from",
			Value: "auths",
		},
			{
				Key:   "localField",
				Value: "auth_id",
			},
			{
				Key:   "foreignField",
				Value: "_id",
			},
			{
				Key:   "as",
				Value: "auth",
			}},
	}}
	unwindStage := bson.D{{Key: "$unwind", Value: bson.D{{Key: "path", Value: "$auth"}, {Key: "preserveNullAndEmptyArrays", Value: false}}}}
	userCollection := db.Collection(UserDB)
	cursor, err := userCollection.Aggregate(ctx, mongo.Pipeline{lookup, unwindStage})
	fmt.Println(cursor)
	var tt []bson.M
	if err = cursor.All(ctx, &tt); err != nil {
		fmt.Println(err)
		return res, err
	}
	data, err := bson.Marshal(&map[string][]bson.M{"values": tt})
	if err != nil {
		fmt.Printf("52\n%v", err)
		return res, err
	}
	var resTest = map[string][]user.UserAggregate{}
	err = bson.Unmarshal(data, &resTest)
	if err != nil {
		fmt.Printf("57\n%v", err)
		return res, err
	}
	fmt.Println(resTest)
	fmt.Println(tt)
	return res, err
}
func (Utb *UserDb) UpdateUser(ctx context.Context, us *user.User) (*user.User, error) {
	// us.Name = "rihisisJHA"
	// mapOfStat := map[string]*user.Status{
	// 	"kkakakaka": &user.Status{
	// 		MarkedAs:  "jab",
	// 		Deadline:  "hja",
	// 		AccessKey: []byte("key"),
	// 	},
	// }
	// us.Ideas = mapOfStat
	db := ctx.Value("mgClient").(*mongo.Database)
	_, err := db.Collection(UserDB).
		ReplaceOne(ctx, bson.D{{"_id", us.ID}},
			*us)

	return us, err
}
