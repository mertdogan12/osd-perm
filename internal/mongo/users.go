package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	Id          int      `bson:"id"`
	Name        string   `bson:"name"`
	Groups      []string `bson:"groups"`
	Permissions []string
}

func GetUser(id int) *User {
	coll := _database.Collection("users")

	var user User
	err := coll.FindOne(context.TODO(), bson.D{{"id", id}}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil
		}

		panic(err)
	}

	return &user
}
