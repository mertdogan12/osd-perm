package mongo

import (
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Group struct {
	Name        string   `bson:"name"`
	Permissions []string `bson:"permissions"`
}

func GetGroup(name string) *Group {
	if _database == nil {
		panic(errors.New("You are not connected (use mongo.Connect() to Connect))"))
	}

	coll := _database.Collection("groups")

	var group Group
	err := coll.FindOne(context.TODO(), bson.D{{"name", name}}).Decode(&group)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			fmt.Println("Couldn't find group:", name)
			return nil
		}

		panic(err)
	}

	return &group
}
