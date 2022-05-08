package mongo

import (
	"context"
	"errors"
	"fmt"

	"github.com/mertdogan12/osd-perm/pkg/helper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	Id          int      `bson:"id" json:"id"`
	Name        string   `bson:"name" json:"name"`
	Groups      []string `bson:"groups" json:"groups"`
	Permissions []string
}

func GetUser(id int) (*User, error) {
	if _database == nil {
		panic(errors.New("You are not connected (use mongo.Connect() to Connect))"))
	}

	coll := _database.Collection("users")

	var user User
	err := coll.FindOne(context.TODO(), bson.D{{"id", id}}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			fmt.Println("Couldn't find user with id:", id)
			return nil, nil
		}

		return nil, err
	}

	perms := make([]string, 0)
	for _, group := range user.Groups {
		g, err := GetGroup(group)

		if err != nil {
			return nil, err
		}

		for _, perm := range g.Permissions {
			if !helper.StringArrayConatins(perms, perm) {
				perms = append(perms, perm)
			}
		}
	}

	user.Permissions = perms

	return &user, nil
}
