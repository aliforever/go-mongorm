package main

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `bson:"_id"`
	FullName string             `bson:"full_name"`
	Username string             `bson:"username"`
}

func (User) Collection() string {
	return "user"
}
